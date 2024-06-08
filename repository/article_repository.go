package repository

import (
	"github.com/C241-PS120/bangkit-cloud-computing/model"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	GetArticleDetail(id int, article *model.Article) error
	GetArticleList(articles *[]model.Article) error
	CreateOrUpdateArticle(article *model.Article) error
	DeleteArticle(id int) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db}
}

func (r *articleRepository) GetArticleDetail(id int, article *model.Article) error {
	return r.db.Preload("Symptoms").
		Preload("Preventions").
		Preload("Treatments").
		Preload("Category").
		Take(&article, id).Error
}

func (r *articleRepository) GetArticleList(articles *[]model.Article) error {
	return r.db.Select("ArticleID", "Title", "ImageURL", "Content").
		Joins("Category").
		Find(&articles).Error
}

func (r *articleRepository) CreateOrUpdateArticle(article *model.Article) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		var category model.Category
		if err := tx.Where("category_name = ?", article.Category.CategoryName).FirstOrCreate(&category).Error; err != nil {
			return err
		}
		article.Category = category

		for i := range article.Symptoms {
			article.Symptoms[i].ArticleID = article.ArticleID
		}
		for i := range article.Preventions {
			article.Preventions[i].ArticleID = article.ArticleID
		}
		for i := range article.Treatments {
			article.Treatments[i].ArticleID = article.ArticleID
		}

		if err := tx.Save(&article).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}

func (r *articleRepository) DeleteArticle(id int) error {
	return r.db.Delete(&model.Article{}, id).Error
}
