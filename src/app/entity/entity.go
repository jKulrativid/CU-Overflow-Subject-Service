package entity

type Post struct {
	Id int64
}

type File struct {
	Id int64
}

type Instructor struct {
	Id             int64
	FullName       string
	Faculty        string
	Email          string
	PhoneNumber    string
	Website        string
	Degree         string
	TaughtSubjects []Subject
}

type Subject struct {
	Id            int64
	SubjectId     string
	Name          string
	Semester      int64
	Section       int64
	Year          int64
	Faculty       string
	Description   string
	Prerequisites []Subject
	Instructors   []Instructor
	Posts         []Post
	Files         []File
}

type PaginationMetadata struct {
	Page       int64
	PerPage    int64
	PageCount  int64
	TotalCount int64
}