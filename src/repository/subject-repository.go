package repository

import (
	"errors"
	"fmt"

	"github.com/jKulrativid/SA-Subject-Service/src/app/entity"
	"gorm.io/gorm"
)

type SubjectRepository struct {
	db      *gorm.DB
	perPage int64
}

func NewSubjectRepository(db *gorm.DB, perPage int64) *SubjectRepository {
	return &SubjectRepository{db: db, perPage: perPage}
}

func (r *SubjectRepository) PaginateSubjects(pageNumber int64, query map[string]interface{}) (*entity.PaginationMetadata, []*entity.Subject, error) {
	subjectId, hasSubjectId := query["subject_id"].(string)
	name, hasName := query["name"].(string)
	semesterWhitelist, hasSemesterWhitelist := query["semester_whitelist"].([]int64)
	sectionWhitelist, hasSectionWhitelist := query["section_whitelist"].([]int64)
	yearRangeStart, hasYearRangeStart := query["year_range_start"].(int64)
	yearRangeStop, hasYearRangeStop := query["year_range_stop"].(int64)

	var subjectRecords []SubjectSchema
	var subjectCount int64

	offset, limit := int((pageNumber-1)*r.perPage), int(r.perPage)

	txErr := r.db.Transaction(func(tx *gorm.DB) error {
		tx = tx.Model(&SubjectSchema{})

		if hasSubjectId {
			tx = tx.Where("subject_id = ?", subjectId)
		}
		if hasName {
			tx = tx.Where("name LIKE ?", fmt.Sprintf("%s%%", name))
		}
		if hasSemesterWhitelist {
			tx = tx.Where("semester IN ?", semesterWhitelist)
		}
		if hasSectionWhitelist {
			tx = tx.Where("section IN ?", sectionWhitelist)
		}
		if hasYearRangeStart && hasYearRangeStop {
			tx = tx.Where("year >= ? AND year <= ?", yearRangeStart, yearRangeStop)
		}

		if err := tx.Count(&subjectCount).Error; err != nil {
			return err
		}

		if err := tx.Offset(offset).Limit(limit).Find(&subjectRecords).Error; err != nil {
			return err
		}

		return nil
	})

	if txErr != nil {
		return nil, nil, txErr
	}

	subjects := make([]*entity.Subject, 0)
	for _, subjectRaw := range subjectRecords {
		subjects = append(subjects, &entity.Subject{
			Id:        int64(subjectRaw.ID),
			SubjectId: subjectRaw.SubjectId,
			Name:      subjectRaw.Name,
			Semester:  subjectRaw.Semester,
			Section:   subjectRaw.Section,
			Year:      subjectRaw.Year,
		})
	}

	pageCount := (subjectCount + r.perPage - 1) / r.perPage

	return &entity.PaginationMetadata{
		Page:       pageNumber,
		PerPage:    r.perPage,
		PageCount:  pageCount,
		TotalCount: subjectCount,
	}, subjects, nil
}

func (r *SubjectRepository) FindSubjectById(id int64) (*entity.Subject, error) {
	var subjectRecord SubjectSchema

	txErr := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Preload("Prerequisites").Preload("Instructors").First(&subjectRecord, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return entity.ErrNotFound
			}
			return err
		}

		return nil
	})

	if txErr != nil {
		return nil, txErr
	}

	subject := entity.Subject{
		Id:            int64(subjectRecord.ID),
		SubjectId:     subjectRecord.SubjectId,
		Name:          subjectRecord.Name,
		Semester:      subjectRecord.Semester,
		Section:       subjectRecord.Section,
		Year:          subjectRecord.Year,
		Faculty:       subjectRecord.Faculty,
		Description:   subjectRecord.Description,
		Prerequisites: make([]entity.Subject, 0),
		Instructors:   make([]entity.Instructor, 0),
	}

	for _, prerequisiteRecord := range subjectRecord.Prerequisites {
		subject.Prerequisites = append(subject.Prerequisites, entity.Subject{
			Id:        int64(prerequisiteRecord.ID),
			SubjectId: prerequisiteRecord.SubjectId,
			Name:      prerequisiteRecord.Name,
		})
	}

	for _, instructorRecord := range subjectRecord.Instructors {
		subject.Instructors = append(subject.Instructors, entity.Instructor{
			Id:       int64(instructorRecord.ID),
			FullName: instructorRecord.FullName,
		})
	}

	return &subject, nil
}

func (r *SubjectRepository) CreateSubject(subject *entity.Subject) error {
	prerequisiteIds := make([]int64, 0)
	prerequisiteRecords := make([]SubjectSchema, 0)
	for _, prerequisite := range subject.Prerequisites {
		prerequisiteIds = append(prerequisiteIds, prerequisite.Id)
		prerequisiteRecords = append(prerequisiteRecords, SubjectSchema{
			Model: gorm.Model{ID: uint(prerequisite.Id)},
		})
	}

	instructorIds := make([]int64, 0)
	instructorRecords := make([]InstructorSchema, 0)
	for _, instructor := range subject.Instructors {
		instructorIds = append(instructorIds, instructor.Id)
		instructorRecords = append(instructorRecords, InstructorSchema{
			Model: gorm.Model{ID: uint(instructor.Id)},
		})
	}

	fmt.Println(prerequisiteRecords)

	subjectRecord := SubjectSchema{
		SubjectId:     subject.SubjectId,
		Name:          subject.Name,
		Semester:      subject.Semester,
		Section:       subject.Section,
		Year:          subject.Year,
		Faculty:       subject.Faculty,
		Description:   subject.Description,
		Prerequisites: prerequisiteRecords,
		Instructors:   instructorRecords,
	}

	txErr := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&subjectRecord).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return entity.ErrConstraintViolation
			}
			return err
		}

		tx.Model(&SubjectSchema{}).Association("Prerequisites").Append(prerequisiteRecords)
		tx.Model(&SubjectSchema{}).Association("Instructors").Append(instructorRecords)

		if len(prerequisiteIds) > 0 {
			if err := tx.Find(&prerequisiteRecords, prerequisiteIds).Error; err != nil {
				return err
			}
		}

		if len(instructorIds) > 0 {
			if err := tx.Find(&instructorRecords, instructorIds).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if txErr != nil {
		return txErr
	}

	subject.Id = int64(subjectRecord.ID)

	prerequisites := make([]entity.Subject, 0)
	for _, prerequisite := range prerequisiteRecords {
		prerequisites = append(prerequisites, entity.Subject{
			Id:        int64(prerequisite.ID),
			SubjectId: prerequisite.SubjectId,
			Name:      prerequisite.Name,
			Faculty:   prerequisite.Faculty,
		})
	}
	subject.Prerequisites = prerequisites

	instructors := make([]entity.Instructor, 0)
	for _, instructor := range instructorRecords {
		instructors = append(instructors, entity.Instructor{
			Id:       int64(instructor.ID),
			FullName: instructor.FullName,
		})
	}
	subject.Instructors = instructors

	return nil
}

func (r *SubjectRepository) UpdateSubject(subject *entity.Subject) error {
	return nil
}

func (r *SubjectRepository) DeleteSubjectById(id int64) error {
	return nil
}
