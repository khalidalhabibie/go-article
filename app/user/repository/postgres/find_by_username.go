package postgres

import (
	"backend/app/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func (r *Repository) FindByUsername(username string) (*models.User, error) {
	result := &models.User{}

	err := r.db.Debug().
	    Where("username = ? ", username).
		First(result).Error
	if err != nil {
		log.Println("error-find-user-by-username: ", err)
		err := fiber.ErrNotFound

		return nil, err
	}

	return result, nil
}
