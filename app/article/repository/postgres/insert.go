package postgres

import (
	"backend/app/models"
	"log"

	"gorm.io/gorm"
)

func (r *Repository) Insert(user *models.Article, tx *gorm.DB) error {
	var db = r.db
	if tx != nil {
		db = tx
	}
	err := db.Create(user).Error
	if err != nil {
		log.Println("error-create-insert: ", err)
		return err
	}
	return nil
}
