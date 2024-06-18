package repository

import (
	"context"
	"fmt"
	"github.com/C241-PS120/bangkit-cloud-computing/model"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	GetArticleById(ctx context.Context, id int, article *model.Article) error
	GetArticleByLabel(ctx context.Context, label string, article *model.Article) error
	GetArticleList(ctx context.Context, articles *[]model.Article) error
	CreateArticle(ctx context.Context, article *model.Article) error
	UpdateArticle(ctx context.Context, article *model.Article) error
	DeleteArticle(ctx context.Context, id int) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db}
}

func (r *articleRepository) GetArticleById(ctx context.Context, id int, article *model.Article) error {
	return r.db.WithContext(ctx).
		Preload("Symptoms").
		Preload("Preventions").
		Preload("Treatments").
		Preload("Disease").
		Preload("Disease.Plants").
		Preload("Label").
		Take(article, id).Error
}

func (r *articleRepository) GetArticleByLabel(ctx context.Context, label string, article *model.Article) error {
	return r.db.WithContext(ctx).Model(&model.Article{}).
		Joins("JOIN label ON article.label_id = label.label_id").
		Preload("Symptoms").
		Preload("Preventions").
		Preload("Treatments").
		Preload("Disease").
		Preload("Disease.Plants").
		Preload("Label").
		Where("label.label_name = ?", label).
		Take(article).Error
}

func (r *articleRepository) GetArticleList(ctx context.Context, articles *[]model.Article) error {
	return r.db.WithContext(ctx).Model(&model.Article{}).
		Preload("Disease", func(db *gorm.DB) *gorm.DB {
			return db.Select("DiseaseID", "DiseaseName")

		}).
		Select("ArticleID", "Title", "ImageURL", "Content", "CreatedAt", "UpdatedAt").
		Find(&articles).Error
}

func (r *articleRepository) CreateArticle(ctx context.Context, article *model.Article) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		if err := r.findOrCreatePlants(tx, &article.Disease.Plants); err != nil {
			return err
		}

		if err := r.findOrCreateDisease(tx, &article.Disease); err != nil {
			return err
		}

		article.DiseaseID = article.Disease.DiseaseID
		var existingArticle model.Article
		if err := tx.Where("disease_id = ?", article.DiseaseID).First(&existingArticle).Error; err == nil {
			return fmt.Errorf("article with the same disease already exists")
		}

		if article.Label.LabelName != "" {
			if err := r.findOrCreateLabel(tx, &article.Label); err != nil {
				return err
			}
			article.LabelID = article.Label.LabelID
		}

		if err := tx.Save(article).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}

func (r *articleRepository) UpdateArticle(ctx context.Context, article *model.Article) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		if err := r.findOrCreateDisease(tx, &article.Disease); err != nil {
			return err
		}
		article.DiseaseID = article.Disease.DiseaseID

		if article.Label.LabelName != "" {
			if err := r.findOrCreateLabel(tx, &article.Label); err != nil {
				return err
			}
			article.LabelID = article.Label.LabelID
		}

		if err := r.findOrCreatePlants(tx, &article.Disease.Plants); err != nil {
			return err
		}

		if err := tx.Model(&article.Disease).Association("Plants").Replace(article.Disease.Plants); err != nil {
			return err
		}

		if err := r.replaceAssociations(tx, article); err != nil {
			return err
		}

		if err := tx.Save(article).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}

func (r *articleRepository) DeleteArticle(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var article model.Article

		// Find the article along with its associations
		if err := tx.Preload("Disease.Plants").Preload("Label").First(&article, id).Error; err != nil {
			return err
		}

		// Clear the association between disease and plants
		if err := tx.Model(&article.Disease).Association("Plants").Clear(); err != nil {
			return err
		}

		// Delete associated symptoms, preventions, and treatments
		if err := tx.Where("article_id = ?", id).Delete(&model.Symptom{}).Error; err != nil {
			return err
		}
		if err := tx.Where("article_id = ?", id).Delete(&model.Prevention{}).Error; err != nil {
			return err
		}
		if err := tx.Where("article_id = ?", id).Delete(&model.Treatment{}).Error; err != nil {
			return err
		}

		// Delete the article
		if err := tx.Delete(&article).Error; err != nil {
			return err
		}

		// Delete the disease if it exists and has no other references
		if article.DiseaseID != 0 {
			var diseaseCount int64
			if err := tx.Model(&model.Article{}).Where("disease_id = ?", article.DiseaseID).Count(&diseaseCount).Error; err != nil {
				return err
			}
			if diseaseCount == 0 {
				if err := tx.Delete(&model.Disease{}, article.DiseaseID).Error; err != nil {
					return err
				}
			}
		}

		// Delete the label if it exists and has no other references
		if article.LabelID != 0 {
			var labelCount int64
			if err := tx.Model(&model.Article{}).Where("label_id = ?", article.LabelID).Count(&labelCount).Error; err != nil {
				return err
			}
			if labelCount == 0 {
				if err := tx.Delete(&model.Label{}, article.LabelID).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})
}

// helper functions
func (r *articleRepository) findOrCreateDisease(tx *gorm.DB, disease *model.Disease) error {
	return tx.Where(model.Disease{DiseaseName: disease.DiseaseName}).
		Attrs(model.Disease{Cause: disease.Cause}).
		FirstOrCreate(disease).Error
}

func (r *articleRepository) findOrCreateLabel(tx *gorm.DB, label *model.Label) error {
	return tx.Where(model.Label{LabelName: label.LabelName}).FirstOrCreate(label).Error
}

func (r *articleRepository) findOrCreatePlants(tx *gorm.DB, plants *[]model.Plant) error {
	for i := range *plants {
		var plant model.Plant
		if err := tx.Where("plant_name = ?", (*plants)[i].PlantName).First(&plant).Error; err != nil {
			if err := tx.Create(&(*plants)[i]).Error; err != nil {
				return err
			}
		} else {
			(*plants)[i] = plant
		}
	}
	return nil
}

func (r *articleRepository) replaceAssociations(tx *gorm.DB, article *model.Article) error {
	var newSymptoms []model.Symptom
	for _, symptom := range article.Symptoms {
		newSymptoms = append(newSymptoms, model.Symptom{SymptomDescription: symptom.SymptomDescription})
	}
	if err := tx.Model(&article).Association("Symptoms").Replace(newSymptoms); err != nil {
		return err
	}

	var newPreventions []model.Prevention
	for _, prevention := range article.Preventions {
		newPreventions = append(newPreventions, model.Prevention{PreventionDescription: prevention.PreventionDescription})
	}
	if err := tx.Model(&article).Association("Preventions").Replace(newPreventions); err != nil {
		return err
	}

	var newTreatments []model.Treatment
	for _, treatment := range article.Treatments {
		newTreatments = append(newTreatments, model.Treatment{
			TreatmentType:        treatment.TreatmentType,
			TreatmentDescription: treatment.TreatmentDescription,
		})
	}
	if err := tx.Model(&article).Association("Treatments").Replace(newTreatments); err != nil {
		return err
	}

	return nil
}
