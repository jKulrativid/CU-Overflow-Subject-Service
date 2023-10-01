package entity

type Post struct {
	Id int64
}

type File struct {
	Id int64
}

type Instructor struct {
	Id            int64
	FullName      string
	Faculty       string
	Email         string
	PhoneNumber   string
	Website       string
	Degree        string
	TaughtSection []Section
}

type Section struct {
	Id          int64
	Number      int64
	Description string
	Instructors []Instructor
	Files       []File
}

type Subject struct {
	Id            int64
	SubjectId     string
	Name          string
	Semester      int64
	Year          int64
	Sections      []Section
	Description   string
	Prerequisites []string
	Faculty       string
	Posts         []Post
}

type SubjectMetadata struct {
	Id        int64
	SubjectId string
	Name      string
	Semester  int64
	Year      int64
}

type InstructorMetadata struct {
	Id       int64
	FullName string
}

type PaginationMetadata struct {
	Page       int64
	PerPage    int64
	PageCount  int64
	TotalCount int64
}
