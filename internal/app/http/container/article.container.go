package container

import (
	"github.com/Article/article-service/internal/article/infrastructure/repository"
	"github.com/Article/article-service/internal/article/transport/http/handler"
	"github.com/Article/article-service/internal/article/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// ArticleContainer holds all dependencies needed for the article HTTP handlers
type ArticleContainer struct {
	Handler *handler.ArticleHandler
}

func NewArticleContainer(db *gorm.DB, logger *logrus.Logger, validator *validator.Validate) *ArticleContainer {
	articleReaderRepository := repository.NewArticleReaderRepository(db)
	articleWriterRepository := repository.NewArticleWriterRepository(db)
	getArticleUsecase := usecase.NewGetArticleUsecase(articleReaderRepository)
	createArticleUsecase := usecase.NewCreateArticleUsecase(articleWriterRepository)
	articleHandler := handler.NewArticleHandler(createArticleUsecase, getArticleUsecase, logger, validator)

	return &ArticleContainer{
		Handler: articleHandler,
	}
}
