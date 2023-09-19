package service

import (
	"context"

	pb "github.com/jKulrativid/SA-Subject-Service/src/grpc/subject"
)

type InstructorService struct {
	pb.UnimplementedInstructorServiceServer
}

func NewInstructorService() *InstructorService {
	return &InstructorService{}
}

func (s *InstructorService) PaginateInstructors(ctx context.Context, req *pb.PaginateInstructorRequest) (*pb.PaginateInstructorResponse, error) {
	return nil, nil
}

func (s *InstructorService) GetInstructorById(ctx context.Context, req *pb.GetInstructorbyIdRequest) (*pb.GetInstructorbyIdResponse, error) {
	return nil, nil
}

func (s *InstructorService) CreateInstructor(ctx context.Context, req *pb.CreateInstructorRequest) (*pb.CreateInstructorResponse, error) {
	return nil, nil
}
