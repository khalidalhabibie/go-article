package postgres

import (
	"backend/app/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (r *Repository) Update(articleM models.Article, tx *gorm.DB) error {
	var db = r.db
	if tx != nil {
		db = tx
	}
	err := db.Save(articleM).Error
	if err != nil {
		log.Println("error-update-article:", err)
		return fiber.ErrUnprocessableEntity
	}
	return nil
}
