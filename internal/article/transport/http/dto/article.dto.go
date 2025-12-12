package dto

import (
	"github.com/google/uuid"
)

type CreateArticleRequest struct {
	Title string `json:"title" validate:"required"`
}

type CreateArticleResponse struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	CreatedAt int64     `json:"created_at"`
}

type GetArticleResponse struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	CreatedAt int64     `json:"created_at"`
}
