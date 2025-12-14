package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/Article/article-service/internal/article/domain/entity"
	articledomainerrors "github.com/Article/article-service/internal/article/domain/errors"
	"github.com/Article/article-service/internal/article/infrastructure/repository/mapper"
	"github.com/Article/article-service/pkg/postgres/models"
	"github.com/Article/article-service/pkg/redis"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ArticleReaderRepository struct {
	db     *gorm.DB
	rdb    *redis.Rdb
	mapper *mapper.ArticleMapper
}

func NewArticleReaderRepository(db *gorm.DB, rdb *redis.Rdb) *ArticleReaderRepository {
	return &ArticleReaderRepository{
		db:     db,
		rdb:    rdb,
		mapper: &mapper.ArticleMapper{},
	}
}

func (r *ArticleReaderRepository) Get(articleId uuid.UUID) (*entity.Article, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("article:%s", articleId.String())

	if r.rdb != nil {
		cached, err := r.rdb.Get(ctx, cacheKey).Result()
		if err == nil {
			var articleModel models.Article
			if err := json.Unmarshal([]byte(cached), &articleModel); err == nil {
				return r.mapper.ToEntity(&articleModel), nil
			}
		}
	}

	articleModel := &models.Article{}
	if err := r.db.First(articleModel, "id = ?", articleId.String()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, articledomainerrors.ErrArticleNotFound
		}
		return nil, err
	}

	article := r.mapper.ToEntity(articleModel)

	if r.rdb != nil {
		articleJSON, err := json.Marshal(articleModel)
		if err == nil {
			r.rdb.Set(ctx, cacheKey, articleJSON, 5*time.Minute)
		}
	}

	return article, nil
}
