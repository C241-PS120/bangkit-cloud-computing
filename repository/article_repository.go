package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/C241-PS120/bangkit-cloud-computing/model"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	GetArticleDetail(ctx context.Context, id int, article *model.Article) error
	GetArticleList(ctx context.Context, articles *[]model.Article) error
	CreateOrUpdateArticle(ctx context.Context, article *model.Article) error
	DeleteArticle(ctx context.Context, id int) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db}
}

func (r *articleRepository) GetArticleDetail(ctx context.Context, id int, article *model.Article) error {
	err := r.db.WithContext(ctx).Preload("Symptoms").
		Preload("Preventions").
		Preload("Treatments").
		Preload("Category").
		Take(&article, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("article with id %d not found", id))
	} else {
		return err
	}
}

func (r *articleRepository) GetArticleList(ctx context.Context, articles *[]model.Article) error {
	return r.db.WithContext(ctx).Select("ArticleID", "Title", "ImageURL", "Content").
		Joins("Category").
		Find(&articles).Error
}

func (r *articleRepository) CreateOrUpdateArticle(ctx context.Context, article *model.Article) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
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

func (r *articleRepository) DeleteArticle(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&model.Article{}, id).Error
}
