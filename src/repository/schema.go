package repository

import (
	"github.com/jKulrativid/SA-Subject-Service/src/app/entity"
	"gorm.io/gorm"
)

type InstructorSchema struct {
	gorm.Model
	FullName       string          `gorm:"type:varchar(255);index:ux_subject"`
	Faculty        string          `gorm:"type:varchar(255)"`
	Email          string          `gorm:"type:varchar(110)"`
	PhoneNumber    string          `gorm:"type:varchar(30)"`
	Website        string          `gorm:"type:varchar(255)"`
	Degree         string          `gorm:"type:text"`
	TaughtSubjects []SubjectSchema `gorm:"many2many:subject_instructors"`
}

func NewInstructorSchema(instructor *entity.Instructor) *InstructorSchema {
	taughtSubjectSchemas := make([]SubjectSchema, 0)
	for _, taughtSubject := range instructor.TaughtSubjects {
		taughtSubjectSchemas = append(taughtSubjectSchemas, *NewSubjectSchema(&taughtSubject))
	}

	return &InstructorSchema{
		Model:          gorm.Model{ID: uint(instructor.Id)},
		FullName:       instructor.FullName,
		Faculty:        instructor.Faculty,
		Email:          instructor.Email,
		PhoneNumber:    instructor.PhoneNumber,
		Website:        instructor.Website,
		Degree:         instructor.Degree,
		TaughtSubjects: taughtSubjectSchemas,
	}
}

func (schema *InstructorSchema) ToInstructor() *entity.Instructor {
	taughtSubjects := make([]entity.Subject, 0)
	for _, taughtSubjectSchema := range schema.TaughtSubjects {
		taughtSubjects = append(taughtSubjects, *taughtSubjectSchema.ToSubject())
	}

	return &entity.Instructor{
		Id:             int64(schema.ID),
		FullName:       schema.FullName,
		Faculty:        schema.Faculty,
		Email:          schema.Email,
		PhoneNumber:    schema.PhoneNumber,
		Website:        schema.Website,
		Degree:         schema.Degree,
		TaughtSubjects: taughtSubjects,
	}
}

func (InstructorSchema) TableName() string {
	return "instructors"
}

type SubjectSchema struct {
	gorm.Model
	SubjectId     string
	Name          string             `gorm:"type:varchar(255);index:ix_subject_name;index:ux_subject_name_semester_section_year_constraint,unique"`
	Semester      int64              `gorm:"index:ix_subject_semester;index:ux_subject_name_semester_section_year_constraint,unique"`
	Section       int64              `gorm:"index:ix_subject_section;index:ux_subject_name_semester_section_year_constraint,unique"`
	Year          int64              `gorm:"index:ix_subject_year;index:ux_subject_name_semester_section_year_constraint,unique"`
	Faculty       string             `gorm:"type:varchar(255)"`
	Description   string             `gorm:"type:text"`
	Prerequisites []SubjectSchema    `gorm:"many2many:subject_prerequisites"`
	Instructors   []InstructorSchema `gorm:"many2many:subject_instructors"`
}

func NewSubjectSchema(subject *entity.Subject) *SubjectSchema {
	prerequisiteSchemas := make([]SubjectSchema, 0)
	for _, prerequisite := range subject.Prerequisites {
		prerequisiteSchemas = append(prerequisiteSchemas, *NewSubjectSchema(&prerequisite))
	}

	instructorSchemas := make([]InstructorSchema, 0)
	for _, instructor := range subject.Instructors {
		instructorSchemas = append(instructorSchemas, *NewInstructorSchema(&instructor))
	}

	return &SubjectSchema{
		Model:         gorm.Model{ID: uint(subject.Id)},
		SubjectId:     subject.SubjectId,
		Name:          subject.Name,
		Semester:      subject.Semester,
		Section:       subject.Section,
		Year:          subject.Year,
		Faculty:       subject.Faculty,
		Description:   subject.Description,
		Prerequisites: prerequisiteSchemas,
		Instructors:   instructorSchemas,
	}
}

func (schema *SubjectSchema) ToSubject() *entity.Subject {
	prerequisites := make([]entity.Subject, 0)
	for _, prerequisiteRecord := range schema.Prerequisites {
		prerequisites = append(prerequisites, *prerequisiteRecord.ToSubject())
	}

	instructors := make([]entity.Instructor, 0)
	for _, instructorRecord := range schema.Instructors {
		instructors = append(instructors, *instructorRecord.ToInstructor())
	}

	return &entity.Subject{
		Id:            int64(schema.ID),
		SubjectId:     schema.SubjectId,
		Name:          schema.Name,
		Semester:      schema.Semester,
		Section:       schema.Section,
		Year:          schema.Year,
		Faculty:       schema.Faculty,
		Description:   schema.Description,
		Prerequisites: prerequisites,
		Instructors:   instructors,
	}
}

func (*SubjectSchema) TableName() string {
	return "subjects"
}
