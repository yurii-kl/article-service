package repository

import (
	"github.com/Article/article-service/internal/article/domain/entity"
	"github.com/Article/article-service/internal/article/infrastructure/repository/mapper"
	"gorm.io/gorm"
)

type ArticleWriterRepository struct {
	db     *gorm.DB
	mapper *mapper.ArticleMapper
}

func NewArticleWriterRepository(db *gorm.DB) *ArticleWriterRepository {
	return &ArticleWriterRepository{db: db, mapper: &mapper.ArticleMapper{}}
}

func (r *ArticleWriterRepository) Create(article *entity.Article) (*entity.Article, error) {
	articleModel := r.mapper.ToModel(article)
	if err := r.db.Create(articleModel).Error; err != nil {
		return nil, err
	}
	return r.mapper.ToEntity(articleModel), nil
}
