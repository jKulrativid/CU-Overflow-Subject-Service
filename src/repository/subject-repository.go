package repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jKulrativid/SA-Subject-Service/src/app/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	for _, subjectRecord := range subjectRecords {
		subjects = append(subjects, subjectRecord.ToSubject())
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
		if err := tx.Preload("Sections").First(&subjectRecord, id).Error; err != nil {
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

	return subjectRecord.ToSubject(), nil
}

func (r *SubjectRepository) CreateSubject(subject *entity.Subject) error {
	subjectRecord := NewSubjectSchema(subject)

	txErr := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit("Sections").Create(subjectRecord).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return entity.ErrConstraintViolation
			}
			return err
		}

		return nil
	})

	if txErr != nil {
		return txErr
	}

	*subject = *subjectRecord.ToSubject()

	return nil
}

func (r *SubjectRepository) UpdateSubject(subject *entity.Subject) error {
	subjectRecord := NewSubjectSchema(subject)

	txErr := r.db.Transaction(func(tx *gorm.DB) error {
		tx = tx.Omit("Sections").Where("id = ?", subject.Id).Updates(subjectRecord)
		if err := tx.Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return entity.ErrConstraintViolation
			}
			return err
		}

		if tx.RowsAffected == 0 {
			return entity.ErrNotFound
		}

		if err := tx.Preload("Sections").First(&subjectRecord, subject.Id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				tx.Rollback()
				return entity.ErrNotFound
			}
			return err
		}

		return nil
	})

	if txErr != nil {
		return txErr
	}

	*subject = *subjectRecord.ToSubject()

	return nil
}

func (r *SubjectRepository) DeleteSubjectById(id int64) (*entity.Subject, error) {
	var subjectRecord SubjectSchema

	txErr := r.db.Transaction(func(tx *gorm.DB) error {
		tx = r.db.Clauses(clause.Returning{}).Delete(&subjectRecord, id)
		if err := tx.Error; err != nil {
			return nil
		}

		if tx.RowsAffected == 0 {
			return entity.ErrNotFound
		}

		return nil
	})

	if txErr != nil {
		return nil, txErr
	}

	return subjectRecord.ToSubject(), nil
}

func (r *SubjectRepository) CreateSection(section *entity.Section) error {
	sectionRecord := NewSectionSchema(section)

	txErr := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit("Instructors.*").Create(sectionRecord).Error; err != nil {
			if strings.Contains(err.Error(), "23503") {
				return entity.ErrNotFound
			}
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return entity.ErrConstraintViolation
			}
			return err
		}

		return nil
	})

	if txErr != nil {
		return txErr
	}

	*section = *sectionRecord.ToSection()

	return nil
}

func (r *SubjectRepository) UpdateSection(section *entity.Section) error {
	return nil
}

func (r *SubjectRepository) DeleteSection(id int64) (*entity.Section, error) {
	return nil, nil
}
