package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/plugin/soft_delete"
)

const (
	ArticleService       = "article"
	ArticleCacheTimeHour = 24
)

type Article struct {
	ID        uuid.UUID `json:"ID" groups:"public"`
	Author    string    `json:"author" groups:"public"`
	Title     string    `json:"title" groups:"public"`
	Body      string    `json:"body" groups:"public"`
	CreatedBy uuid.UUID `json:"created_by" groups:"public"`
	CreatedAt time.Time `json:"created_at" groups:"public"`
	UpdatedAt time.Time `json:"updated_at" `
	DeletedAt soft_delete.DeletedAt
}
