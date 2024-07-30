package database

import (
	"log"

	mocket "github.com/Selvatico/go-mocket"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDBTests() *gorm.DB {
	mocket.Catcher.Register()

	db, err := gorm.Open(postgres.New(postgres.Config{DriverName: mocket.DriverName}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatalf("error mocking gorm: %s", err)
	}

	return db
}
