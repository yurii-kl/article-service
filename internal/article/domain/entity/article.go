package entity

import (
	"github.com/google/uuid"
	"time"
)

type Article struct {
	id        uuid.UUID
	title     string
	createdAt time.Time
}

func NewArticle(title string) *Article {
	return &Article{
		id:        uuid.New(),
		title:     title,
		createdAt: time.Now().UTC(),
	}
}

func NewArticleWithID(id uuid.UUID, title string, createdAt time.Time) *Article {
	return &Article{
		id:        id,
		title:     title,
		createdAt: createdAt,
	}
}

func (a *Article) ID() uuid.UUID        { return a.id }
func (a *Article) Title() string        { return a.title }
func (a *Article) CreatedAt() time.Time { return a.createdAt }
