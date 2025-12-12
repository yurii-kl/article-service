package container

import (
	"github.com/google/uuid"
	"testing"

	"github.com/Article/article-service/internal/article/infrastructure/repository"
	"github.com/Article/article-service/internal/article/usecase"
	"github.com/Article/article-service/pkg/postgres/models"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	err = db.AutoMigrate(&models.Article{})
	if err != nil {
		t.Fatalf("Failed to migrate: %v", err)
	}

	return db
}

func TestNewArticleContainer_InitializesCorrectly(t *testing.T) {
	db := setupTestDB(t)
	logger := logrus.New()
	validator := validator.New()

	container := NewArticleContainer(db, nil, logger, validator)

	assert.NotNil(t, container)
	assert.NotNil(t, container.Handler)
}

func TestArticleContainer_Integration_CreateAndGetArticle(t *testing.T) {
	db := setupTestDB(t)

	articleReaderRepository := repository.NewArticleReaderRepository(db, nil)
	articleWriterRepository := repository.NewArticleWriterRepository(db, nil)
	createUsecase := usecase.NewCreateArticleUsecase(articleWriterRepository)
	getUsecase := usecase.NewGetArticleUsecase(articleReaderRepository)

	title := "Integration Test Article"
	createdArticle, err := createUsecase.Execute(title)

	assert.NoError(t, err)
	assert.NotNil(t, createdArticle)
	assert.Equal(t, title, createdArticle.Title())
	assert.NotEqual(t, uuid.Nil, createdArticle.ID())

	retrievedArticle, err := getUsecase.Execute(createdArticle.ID())

	// Assert - Get
	assert.NoError(t, err)
	assert.NotNil(t, retrievedArticle)
	assert.Equal(t, createdArticle.ID(), retrievedArticle.ID())
	assert.Equal(t, title, retrievedArticle.Title())
}
