package postgres

import (
	"backend/app/models"
	"backend/pkg/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

func (r *Repository) FindAll(config utils.PaginationConfig) ([]models.Article, error) {
	results := []models.Article{}

	err := r.db.Debug().
		Scopes(config.MetaScopes()...).
		Scopes(config.Scopes()...).
		Find(&results).Error
	if err != nil {
		log.Println("error-find-all-article: ", err)
		err := fiber.ErrUnprocessableEntity

		return nil, err
	}

	return results, nil
}
