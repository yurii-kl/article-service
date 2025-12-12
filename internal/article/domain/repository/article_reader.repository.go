package repository

import (
	"github.com/Article/article-service/internal/article/domain/entity"
	"github.com/google/uuid"
)

type ArticleReaderRepository interface {
	Get(articleId uuid.UUID) (*entity.Article, error)
}
