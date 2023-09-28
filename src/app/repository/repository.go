package repository

import "github.com/jKulrativid/SA-Subject-Service/src/app/entity"

type InstructorRepository interface {
}

type SubjectRepository interface {
	PaginateSubjects(pageNumber int64, query map[string]interface{}) (*entity.PaginationMetadata, []*entity.Subject, error)
	FindSubjectById(id int64) (*entity.Subject, error)
	CreateSubject(subject *entity.Subject) error
	UpdateSubject(subject *entity.Subject) error
	DeleteSubjectById(id int64) (*entity.Subject, error)
}
