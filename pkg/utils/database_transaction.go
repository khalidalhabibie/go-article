package utils

import "gorm.io/gorm"

type TransactionManagerClient interface {
	NewTransaction() *gorm.DB
}

type TransactionManager struct {
	db *gorm.DB
}

func TransactionManagerNew(db *gorm.DB) TransactionManagerClient {
	return &TransactionManager{
		db: db,
	}
}

func (tm *TransactionManager) NewTransaction() *gorm.DB {
	return tm.db.Begin()
}

func databaseTransaction(db *gorm.DB) *gorm.DB {

	return db.Begin()

}
