package postgres

import (
	"backend/app/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func (r *Repository) FindByUsernameAndIsVerified(username string, isverified bool) (*models.User, error) {
	result := &models.User{}

	err := r.db.Debug().
		Where("username = ? ", username).
		Where("is_verified = ? ", isverified).
		First(result).Error
	if err != nil {
		log.Println("error-find-user-by-username-and-is-verified: ", err)
		err := fiber.ErrNotFound

		return nil, err
	}

	return result, nil
}
