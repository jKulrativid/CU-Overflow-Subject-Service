package repository

import (
	"github.com/jKulrativid/SA-Subject-Service/src/app/entity"
	"gorm.io/gorm"
)

type InstructorRepository struct {
	db      *gorm.DB
	perPage int64
}

func NewInstructorRepository(db *gorm.DB, perPage int64) *InstructorRepository {
	return &InstructorRepository{db: db, perPage: perPage}
}

func (r *InstructorRepository) CreateInstructor(instructor *entity.Instructor) error {
	return nil
}
