package http

import (
	"github.com/Article/article-service/internal/app/http/container"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(v1 *gin.RouterGroup, articleContainer *container.ArticleContainer, md ...gin.HandlerFunc) {
	article := v1.Group("/article", md...)
	article.POST("/", articleContainer.Handler.CreateArticle)
	article.GET("/:id", articleContainer.Handler.GetArticle)
}
