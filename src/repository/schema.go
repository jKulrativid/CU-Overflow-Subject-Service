package repository

import (
	"github.com/jKulrativid/SA-Subject-Service/src/app/entity"
	"github.com/lib/pq"
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
	TaughtSections []SectionSchema `gorm:"many2many:subject_instructors"`
}

func NewInstructorSchema(instructor *entity.Instructor) *InstructorSchema {
	taughtSectionSchemas := make([]SectionSchema, 0)
	for _, taughtSubject := range instructor.TaughtSection {
		taughtSectionSchemas = append(taughtSectionSchemas, *NewSectionSchema(&taughtSubject))
	}

	return &InstructorSchema{
		Model:          gorm.Model{ID: uint(instructor.Id)},
		FullName:       instructor.FullName,
		Faculty:        instructor.Faculty,
		Email:          instructor.Email,
		PhoneNumber:    instructor.PhoneNumber,
		Website:        instructor.Website,
		Degree:         instructor.Degree,
		TaughtSections: taughtSectionSchemas,
	}
}

func (schema *InstructorSchema) ToInstructor() *entity.Instructor {
	taughtSection := make([]entity.Section, 0)
	for _, taughtSectionSchema := range schema.TaughtSections {
		taughtSection = append(taughtSection, *taughtSectionSchema.ToSection())
	}

	return &entity.Instructor{
		Id:            int64(schema.ID),
		FullName:      schema.FullName,
		Faculty:       schema.Faculty,
		Email:         schema.Email,
		PhoneNumber:   schema.PhoneNumber,
		Website:       schema.Website,
		Degree:        schema.Degree,
		TaughtSection: taughtSection,
	}
}

func (InstructorSchema) TableName() string {
	return "instructors"
}

type SectionSchema struct {
	gorm.Model
	SubjectId   uint               `gorm:"index:ux_subject_id_section_number,unique"`
	Number      int64              `gorm:"not null;index:ux_subject_id_section_number,unique"`
	Description string             `gorm:"type:text"`
	Instructors []InstructorSchema `gorm:"many2many:section_instructors"`
}

func NewSectionSchema(section *entity.Section) *SectionSchema {
	instructorSchemas := make([]InstructorSchema, 0)
	for _, instructor := range section.Instructors {
		instructorSchemas = append(instructorSchemas, *NewInstructorSchema(&instructor))
	}

	return &SectionSchema{
		Model:       gorm.Model{ID: uint(section.Id)},
		SubjectId:   uint(section.SubjectId),
		Number:      section.Number,
		Description: section.Description,
		Instructors: instructorSchemas,
	}
}

func (schema *SectionSchema) ToSection() *entity.Section {
	instructors := make([]entity.Instructor, 0)
	for _, instructorSchema := range schema.Instructors {
		instructors = append(instructors, *instructorSchema.ToInstructor())
	}
	return &entity.Section{
		Id:          int64(schema.ID),
		SubjectId:   int64(schema.SubjectId),
		Number:      schema.Number,
		Description: schema.Description,
		Instructors: instructors,
	}
}

func (*SectionSchema) TableName() string {
	return "sections"
}

type SubjectSchema struct {
	gorm.Model
	SubjectId     string
	Name          string          `gorm:"not null;type:varchar(255);index:ix_subject_name;index:ux_subject_name_semester_year_constraint,unique"`
	Semester      int64           `gorm:"not null;index:ix_subject_semester;index:ux_subject_name_semester_year_constraint,unique"`
	Sections      []SectionSchema `gorm:"foreignKey:SubjectId;references:ID"`
	Year          int64           `gorm:"not null;index:ix_subject_year;index:ux_subject_name_semester_year_constraint,unique"`
	Faculty       string          `gorm:"type:varchar(255)"`
	Description   string          `gorm:"type:text"`
	Prerequisites pq.StringArray  `gorm:"type:varchar(64)[]"`
}

func NewSubjectSchema(subject *entity.Subject) *SubjectSchema {
	sectionSchemas := make([]SectionSchema, 0)
	for _, section := range subject.Sections {
		sectionSchemas = append(sectionSchemas, *NewSectionSchema(&section))
	}

	return &SubjectSchema{
		Model:         gorm.Model{ID: uint(subject.Id)},
		SubjectId:     subject.SubjectId,
		Name:          subject.Name,
		Semester:      subject.Semester,
		Sections:      sectionSchemas,
		Year:          subject.Year,
		Faculty:       subject.Faculty,
		Description:   subject.Description,
		Prerequisites: subject.Prerequisites,
	}
}

func (schema *SubjectSchema) ToSubject() *entity.Subject {
	sections := make([]entity.Section, 0)
	for _, sectionRecord := range schema.Sections {
		sections = append(sections, *sectionRecord.ToSection())
	}

	return &entity.Subject{
		Id:            int64(schema.ID),
		SubjectId:     schema.SubjectId,
		Name:          schema.Name,
		Semester:      schema.Semester,
		Sections:      sections,
		Year:          schema.Year,
		Faculty:       schema.Faculty,
		Description:   schema.Description,
		Prerequisites: schema.Prerequisites,
	}
}

func (*SubjectSchema) TableName() string {
	return "subjects"
}
