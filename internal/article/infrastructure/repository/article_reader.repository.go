package repository

import (
	"github.com/Article/article-service/internal/article/domain/entity"
	"github.com/Article/article-service/internal/article/infrastructure/repository/mapper"
	"github.com/Article/article-service/pkg/postgres/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ArticleReaderRepository struct {
	db     *gorm.DB
	mapper *mapper.ArticleMapper
}

func NewArticleReaderRepository(db *gorm.DB) *ArticleReaderRepository {
	return &ArticleReaderRepository{db: db, mapper: &mapper.ArticleMapper{}}
}

func (r *ArticleReaderRepository) Get(articleId uuid.UUID) (*entity.Article, error) {
	articleModel := &models.Article{}
	if err := r.db.First(articleModel, "id = ?", articleId.String()).Error; err != nil {
		return nil, err
	}
	return r.mapper.ToEntity(articleModel), nil
}
