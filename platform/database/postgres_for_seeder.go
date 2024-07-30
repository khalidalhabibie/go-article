package database

import (
	"backend/pkg/utils"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// func PostgreSQLConnection() (*gorm.DB, error) {
func PostgreSQLSeedConnection() *gorm.DB {
	// Define database connection settings.
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	// Build PostgreSQL connection URL.
	postgresConnURL, err := utils.ConnectionURLBuilder("postgres_seeder")
	if err != nil {
		// return nil, err
		log.Panic("error get data builder, err: ", err)

	}

	db, err := gorm.Open(postgres.Open(postgresConnURL), &gorm.Config{

		Logger: logger.Default.LogMode(logger.Info),
		// Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		// log.Println(err)
		// panic(err)
		// return nil, fmt.Errorf("error, not connected to database, %w", err)
		log.Panic("error  connected to database, err: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		// return nil, fmt.Errorf("error to assign db to sql db ")
		log.Panic("error to assign db to sql db, err: ", err)
	}

	sqlDB.SetMaxOpenConns(maxConn)
	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	if err := sqlDB.Ping(); err != nil {
		defer sqlDB.Close() // close database connection
		// return nil, fmt.Errorf("error, not sent ping to database, %w", err)
		log.Panic("error, not sent ping to database: ", err)
	}

	// return db, nil
	return db

}
