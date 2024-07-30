package postgres

import (
	"backend/app/models"
	"backend/pkg/utils"
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

func (r *Repository) Count(config utils.PaginationConfig) (int64, error) {
	var count int64

	err := r.db.
		Model(&models.Article{}).
		Scopes(config.Scopes()...).
		Count(&count).Error
	if err != nil {
		log.Println("error-count-article:", err)
		err := fiber.ErrUnprocessableEntity
		return 0, err
	}

	return count, nil
}
