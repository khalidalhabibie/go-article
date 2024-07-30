package user

import (
	"backend/app/models"
	"backend/pkg/utils"

	"gorm.io/gorm"
)

type Repository interface {
	FindByUsername(username string) (*models.User, error)
	FindByEmail(username string) (*models.User, error)
	Insert(user models.User, tx *gorm.DB) error
	Update(userM models.User, tx *gorm.DB) error
	FindAll(config utils.PaginationConfig) ([]models.User, error)
	Count(config utils.PaginationConfig) (int64, error)
}

type Cache interface {
	SetVerificationCode(email, verificationCode string) error
	GetVerificationCode(email, verificationCode string) (*string, error)
	DelVerificationCode(email string) error
}
