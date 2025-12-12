package repository

import "github.com/Article/article-service/internal/article/domain/entity"

type ArticleWriterRepository interface {
	Create(article *entity.Article) (*entity.Article, error)
}
