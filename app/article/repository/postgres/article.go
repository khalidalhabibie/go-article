package postgres

import (
	"backend/app/article"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) article.Repository {
	return &Repository{
		db: db,
	}
}
