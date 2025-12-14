package test

import (
	"errors"
	"testing"
	"time"

	"github.com/Article/article-service/internal/article/domain/entity"
	"github.com/Article/article-service/internal/article/infrastructure/repository/mocks"
	"github.com/Article/article-service/internal/article/usecase"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetArticleUsecase_Execute_Success(t *testing.T) {
	mockReader := new(mocks.MockArticleReaderRepository)
	usecase := usecase.NewGetArticleUsecase(mockReader)

	articleID := uuid.New()
	title := "Test Article Title"
	createdAt := time.Now().UTC()
	updatedAt := time.Now().UTC()
	expectedArticle := entity.NewArticleWithID(articleID, title, createdAt, updatedAt)

	mockReader.On("Get", articleID).Return(expectedArticle, nil)

	result, err := usecase.Execute(articleID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, articleID, result.ID())
	assert.Equal(t, title, result.Title())
	mockReader.AssertExpectations(t)
}

func TestGetArticleUsecase_Execute_ArticleNotFound(t *testing.T) {
	mockReader := new(mocks.MockArticleReaderRepository)
	usecase := usecase.NewGetArticleUsecase(mockReader)

	articleID := uuid.New()
	expectedError := errors.New("article not found")

	mockReader.On("Get", articleID).Return(nil, expectedError)

	result, err := usecase.Execute(articleID)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedError, err)
	mockReader.AssertExpectations(t)
}
