package repository

import (
	"github.com/jKulrativid/SA-Subject-Service/src/app/entity"
	"gorm.io/gorm"
)

type SubjectRepository struct {
	db *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) *SubjectRepository {
	return &SubjectRepository{db: db}
}

func (r *SubjectRepository) PaginateSubjects(query map[string]interface{}) (*entity.PaginationMetadata, []*entity.Subject, error) {
	return nil, nil, nil
}

func (r *SubjectRepository) FindSubjectById(id int64) (*entity.Subject, error) {
	return nil, nil
}

func (r *SubjectRepository) CreateSubject(subject *entity.Subject) error {
	return nil
}

func (r *SubjectRepository) UpdateSubject(subject *entity.Subject) error {
	return nil
}

func (r *SubjectRepository) DeleteSubjectById(id int64) error {
	return nil
}
