package handler

import (
	"net/http"

	"github.com/Article/article-service/internal/article/transport/http/dto"
	"github.com/Article/article-service/internal/article/usecase"
	"github.com/Article/article-service/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ArticleHandler struct {
	createArticleUsecase *usecase.CreateArticleUsecase
	getArticleUsecase    *usecase.GetArticleUsecase
	logger               *logrus.Logger
	validator            *validator.Validate
}

func NewArticleHandler(
	createArticleUsecase *usecase.CreateArticleUsecase,
	getArticleUsecase *usecase.GetArticleUsecase,
	logger *logrus.Logger,
	validator *validator.Validate,
) *ArticleHandler {
	return &ArticleHandler{
		createArticleUsecase: createArticleUsecase,
		getArticleUsecase:    getArticleUsecase,
		logger:               logger,
		validator:            validator,
	}
}

func (ah *ArticleHandler) CreateArticle(c *gin.Context) {
	var req dto.CreateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ah.logger.WithError(err).Error("invalid request body")
		response.BadRequestBody(c, err)
		return
	}

	if err := ah.validator.Struct(&req); err != nil {
		ah.logger.WithError(err).Error("validation failed")
		response.BadRequestBody(c, err)
		return
	}

	article, err := ah.createArticleUsecase.Execute(req.Title)
	if err != nil {
		ah.logger.WithError(err).Error("create article failed")
		response.InternalError(c, err)
		return
	}

	resp := dto.CreateArticleResponse{
		ID:        article.ID(),
		Title:     article.Title(),
		CreatedAt: article.CreatedAt().Unix(),
	}
	c.JSON(http.StatusOK, resp)
}

func (ah *ArticleHandler) GetArticle(c *gin.Context) {
	articleIDStr := c.Param("id")
	articleID, err := uuid.Parse(articleIDStr)
	if err != nil {
		ah.logger.WithError(err).Error("invalid request body")
		response.BadRequestPath(c, err)
		return
	}

	article, err := ah.getArticleUsecase.Execute(articleID)
	if err != nil {
		ah.logger.WithError(err).Error("get article failed")
		response.InternalError(c, err)
		return
	}

	resp := dto.GetArticleResponse{
		ID:        article.ID(),
		Title:     article.Title(),
		CreatedAt: article.CreatedAt().Unix(),
	}
	c.JSON(http.StatusOK, resp)
}
