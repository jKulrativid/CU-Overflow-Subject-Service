package service

type Post struct {
}

type File struct {
}

type Instructor struct {
	Id          int64
	FullName    string
	Email       string
	PhoneNumber string
	Website     string
	Degree      string
}

type Subject struct {
	Id            int64
	Name          string
	Semester      int64
	Year          int64
	Description   string
	Prerequisites []Subject
	Instructors   []Instructor
	Posts         []Post
	Files         []File
}
