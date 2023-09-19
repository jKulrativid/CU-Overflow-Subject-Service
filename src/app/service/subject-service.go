package service

import (
	"context"

	pb "github.com/jKulrativid/SA-Subject-Service/src/grpc/subject"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SubjectService struct {
	pb.UnimplementedSubjectServiceServer
}

func NewSubjectService() *SubjectService {
	return &SubjectService{}
}

func (s *SubjectService) PaginateSubjects(ctx context.Context, req *pb.PaginateSubjectRequest) (*pb.PaginateSubjectResponse, error) {
	subjects := make([]*pb.SubjectMetadata, 3)
	for i := 0; i < 3; i++ {
		subjects[i] = &pb.SubjectMetadata{
			Id:        int64(i),
			SubjectId: "2110521",
			Name:      "Software Architecture",
			Semester:  1,
			Section:   int64(i),
			Year:      2023,
		}
	}

	return &pb.PaginateSubjectResponse{
		PageNumber: 1,
		PerPage:    3,
		PageCount:  1,
		TotalCount: 3,
		Subjects:   subjects,
	}, nil
}

func (s *SubjectService) GetSubjectById(ctx context.Context, req *pb.GetSubjectByIdRequest) (*pb.GetSubjectByIdResponse, error) {
	if req.Id < 0 || req.Id > 2 {
		return nil, status.Error(codes.NotFound, "subject with given ID not found")
	}

	id := req.Id

	return &pb.GetSubjectByIdResponse{
		Subject: &pb.Subject{
			Id:          id,
			SubjectId:   "2110521",
			Name:        "Software Architecture",
			Semester:    1,
			Section:     id,
			Year:        2023,
			Faculty:     "Engineering",
			Description: "このサビスはモクトです",
			Instructors: []*pb.Instructor{
				{
					Id:       1,
					FullName: "Assoc. Prof. Kulwadee Something",
					Faculty:  "Engineering",
					Email:    "kulwadee-sds@chula.ac.th",
				},
			},
		},
	}, nil
}

/*
func (s *SubjectService) CreateSubject(ctx context.Context, req *pb.CreateSubjectRequest) (*pb.CreateSubjectResponse, error) {
	return nil, nil
}

func (s *SubjectService) UpdateSubject(ctx context.Context, req *pb.UpdateSubjectRequest) (*pb.UpdateSubjectResponse, error) {
	return nil, nil
}

func (s *SubjectService) DeleteSubject(ctx context.Context, req *pb.DeleteSubjectRequest) (*pb.DeleteSubjectResponse, error) {
	return nil, nil
}
*/
