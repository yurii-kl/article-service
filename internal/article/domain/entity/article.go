package entity

import (
	"github.com/google/uuid"
	"time"
)

type Article struct {
	id        uuid.UUID
	title     string
	createdAt time.Time
	updatedAt time.Time
}

func NewArticle(title string) *Article {
	return &Article{
		id:        uuid.New(),
		title:     title,
		createdAt: time.Now().UTC(),
		updatedAt: time.Now().UTC(),
	}
}

func NewArticleWithID(id uuid.UUID, title string, createdAt time.Time, updatedAt time.Time) *Article {
	return &Article{
		id:        id,
		title:     title,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (a *Article) ID() uuid.UUID        { return a.id }
func (a *Article) Title() string        { return a.title }
func (a *Article) CreatedAt() time.Time { return a.createdAt }
func (a *Article) UpdatedAt() time.Time { return a.updatedAt }
