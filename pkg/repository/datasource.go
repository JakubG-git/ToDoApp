package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDataSource creates a new gorm.DB instance
func NewDataSource(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ConfigureDataSource(dsn string) (*gorm.DB, error) {
	db, err := NewDataSource(dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Migration(db *gorm.DB, dst ...interface{}) error {
	return db.AutoMigrate(dst...)
}
