package mapper

import (
	"github.com/Article/article-service/internal/article/domain/entity"
	"github.com/Article/article-service/pkg/postgres/models"
)

type ArticleMapper struct{}

func (m *ArticleMapper) ToEntity(model *models.Article) *entity.Article {
	return entity.NewArticleWithID(
		model.Id,
		model.Title,
		model.CreatedAt,
	)
}

func (m *ArticleMapper) ToModel(e *entity.Article) *models.Article {
	return &models.Article{
		Id:    e.ID(),
		Title: e.Title(),
	}
}
