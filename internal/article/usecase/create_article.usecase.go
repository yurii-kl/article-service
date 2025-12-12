package usecase

import (
	"github.com/Article/article-service/internal/article/domain/entity"
	"github.com/Article/article-service/internal/article/domain/repository"
)

type CreateArticleUsecase struct {
	articleWriter repository.ArticleWriterRepository
}

func NewCreateArticleUsecase(articleWriter repository.ArticleWriterRepository) *CreateArticleUsecase {
	return &CreateArticleUsecase{
		articleWriter: articleWriter,
	}
}

func (u *CreateArticleUsecase) Execute(title string) (*entity.Article, error) {
	newArticle := entity.NewArticle(title)

	createdArticle, err := u.articleWriter.Create(newArticle)
	if err != nil {
		return nil, err
	}

	return createdArticle, nil
}
