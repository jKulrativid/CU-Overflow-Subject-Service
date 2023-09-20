package repository

import (
	"gorm.io/gorm"
)

type InstructorRepository struct {
	db *gorm.DB
}

func NewInstructorRepository(db *gorm.DB) *InstructorRepository {
	return &InstructorRepository{db: db}
}
