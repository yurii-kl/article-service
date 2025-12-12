package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Id    uuid.UUID
	Title string `gorm:"type:varchar(255);not null"`
}
