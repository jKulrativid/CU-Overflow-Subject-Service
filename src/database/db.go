package database

import (
	"errors"
	"fmt"

	"github.com/jKulrativid/SA-Subject-Service/src/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	Port     int
	SslMode  string
}

func NewDatabaseConnection(config *DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%d sslmode=%s", config.Host, config.User, config.Password, config.Port, config.SslMode)
	db, err := gorm.Open(postgres.New(postgres.Config{DriverName: "pgx", DSN: dsn}), &gorm.Config{TranslateError: true})
	if err != nil {
		return nil, errors.New("failed to connected to DB")
	}
	return db, nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&repository.SubjectSchema{}, &repository.SectionSchema{}, &repository.InstructorSchema{})
}
