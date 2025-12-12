package usecase

import (
	"github.com/Article/article-service/internal/article/domain/entity"
	"github.com/Article/article-service/internal/article/domain/repository"
	"github.com/google/uuid"
)

type GetArticleUsecase struct {
	articleReader repository.ArticleReaderRepository
}

func NewGetArticleUsecase(articleReader repository.ArticleReaderRepository) *GetArticleUsecase {
	return &GetArticleUsecase{
		articleReader: articleReader,
	}
}

func (u *GetArticleUsecase) Execute(articleId uuid.UUID) (*entity.Article, error) {
	return u.articleReader.Get(articleId)
}
