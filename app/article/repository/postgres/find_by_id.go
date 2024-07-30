package postgres

import (
	"backend/app/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (r *Repository) FindById(id uuid.UUID) (*models.Article, error) {
	result := &models.Article{}

	err := r.db.Debug().
		Where("id = ? ", id).
		First(result).Error
	if err != nil {
		log.Println("error-find-article-by-id: ", err)
		err := fiber.ErrNotFound

		return nil, err
	}

	return result, nil
}
