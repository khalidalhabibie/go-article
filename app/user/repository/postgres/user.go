package postgres

import (
	"backend/app/user"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.Repository {
	return &Repository{
		db: db,
	}
}
