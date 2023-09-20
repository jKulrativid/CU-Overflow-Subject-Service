package repository

import "gorm.io/gorm"

type SubjectSchema struct {
	gorm.Model
	SubjectId     string
	Name          string          `gorm:"type:varchar(255);index:ix_subject_name;index:ux_subject_name_semester_section_year_constraint,unique"`
	Semester      int64           `gorm:"index:ix_subject_semester;index:ux_subject_name_semester_section_year_constraint,unique"`
	Section       int64           `gorm:"index:ix_subject_section;index:ux_subject_name_semester_section_year_constraint,unique"`
	Year          int64           `gorm:"index:ix_subject_year;index:ux_subject_name_semester_section_year_constraint,unique"`
	Faculty       string          `gorm:"type:varchar(255)"`
	Description   string          `gorm:"type:text"`
	Prerequisites []SubjectSchema `gorm:"foreignkey:ID"`
}

func (SubjectSchema) TableName() string {
	return "subjects"
}
