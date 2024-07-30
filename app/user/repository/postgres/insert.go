package postgres

import (
	"backend/app/models"
	"log"

	"gorm.io/gorm"
)

func (r *Repository) Insert(user models.User, tx *gorm.DB) error {
	var db = r.db
	if tx != nil {
		db = tx
	}

	log.Println("user ", user)
	err := db.Debug().Create(user).Error
	if err != nil {
		log.Println("error-insert-user: ", err)
		return err
	}
	return nil
}
