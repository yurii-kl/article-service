package test

import (
	"errors"
	"testing"

	"github.com/Article/article-service/internal/article/domain/entity"
	"github.com/Article/article-service/internal/article/infrastructure/repository/mocks"
	"github.com/Article/article-service/internal/article/usecase"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateArticleUsecase_Execute_Success(t *testing.T) {
	mockWriter := new(mocks.MockArticleWriterRepository)
	usecase := usecase.NewCreateArticleUsecase(mockWriter)

	title := "Test Article Title"
	expectedArticle := entity.NewArticle(title)

	mockWriter.On("Create", mock.AnythingOfType("*entity.Article")).Return(expectedArticle, nil)

	result, err := usecase.Execute(title)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, title, result.Title())
	assert.NotEqual(t, uuid.Nil, result.ID())
	mockWriter.AssertExpectations(t)
}

func TestCreateArticleUsecase_Execute_RepositoryError(t *testing.T) {
	mockWriter := new(mocks.MockArticleWriterRepository)
	usecase := usecase.NewCreateArticleUsecase(mockWriter)

	title := "Test Article Title"
	expectedError := errors.New("database connection failed")

	mockWriter.On("Create", mock.AnythingOfType("*entity.Article")).Return(nil, expectedError)

	result, err := usecase.Execute(title)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedError, err)
	mockWriter.AssertExpectations(t)
}
