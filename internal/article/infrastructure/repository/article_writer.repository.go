package repository

import (
	"context"
	"fmt"

	"github.com/Article/article-service/internal/article/domain/entity"
	"github.com/Article/article-service/internal/article/infrastructure/repository/mapper"
	"github.com/Article/article-service/pkg/redis"
	"gorm.io/gorm"
)

type ArticleWriterRepository struct {
	db     *gorm.DB
	rdb    *redis.Rdb
	mapper *mapper.ArticleMapper
}

func NewArticleWriterRepository(db *gorm.DB, rdb *redis.Rdb) *ArticleWriterRepository {
	return &ArticleWriterRepository{db: db, rdb: rdb, mapper: &mapper.ArticleMapper{}}
}

func (r *ArticleWriterRepository) Create(article *entity.Article) (*entity.Article, error) {
	articleModel := r.mapper.ToModel(article)
	if err := r.db.Create(articleModel).Error; err != nil {
		return nil, err
	}

	createdArticle := r.mapper.ToEntity(articleModel)

	if r.rdb != nil {
		cacheKey := fmt.Sprintf("article:%s", createdArticle.ID().String())
		r.rdb.Del(context.Background(), cacheKey)
	}

	return createdArticle, nil
}
