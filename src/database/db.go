package database

import (
	"errors"

	"github.com/jKulrativid/SA-Subject-Service/src/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection() (*gorm.DB, error) {
	dsn := "host=subject-db user=subject password=please_use_long_passphrase_ibeg_you port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to connected to DB")
	}
	return db, nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&repository.SubjectSchema{})
}
