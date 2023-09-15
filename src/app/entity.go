package service

type Post struct {
}

type File struct {
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
	Year          int64
	Faculty       string
	Description   string
	Prerequisites []Subject
	Instructors   []Instructor
	Posts         []Post
	Files         []File
}
