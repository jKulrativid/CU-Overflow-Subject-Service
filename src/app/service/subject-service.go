package service

import (
	"context"
	"fmt"

	"github.com/jKulrativid/SA-Subject-Service/src/app/entity"
	"github.com/jKulrativid/SA-Subject-Service/src/app/repository"
	pb "github.com/jKulrativid/SA-Subject-Service/src/grpc/subject"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SubjectService struct {
	pb.UnimplementedSubjectServiceServer
	subjectRepo repository.SubjectRepository
}

func NewSubjectService(subjectRepo repository.SubjectRepository) *SubjectService {
	return &SubjectService{subjectRepo: subjectRepo}
}

func (s *SubjectService) PaginateSubjects(ctx context.Context, req *pb.PaginateSubjectRequest) (*pb.PaginateSubjectResponse, error) {
	if req.PageNumber < 1 {
		return nil, status.Error(codes.InvalidArgument, "page number must be a positive integer")
	}

	query := make(map[string]interface{})

	if req.SubjectId != "" {
		query["subject_id"] = req.SubjectId
	}
	if req.Name != "" {
		query["name"] = req.Name
	}
	if len(req.SemesterWhitelist) != 0 {
		query["semester_whitelist"] = req.SemesterWhitelist
	}
	if len(req.SectionWhitelist) != 0 {
		query["section_whitelist"] = req.SectionWhitelist
	}
	if req.YearRangeStart != 0 && req.YearRangeStop != 0 {
		query["year_range_start"] = req.YearRangeStart
		query["year_range_stop"] = req.YearRangeStop
	}

	metadata, subjects, err := s.subjectRepo.PaginateSubjects(req.PageNumber, query)
	if err != nil {
		return nil, err
	}

	resp := pb.PaginateSubjectResponse{
		PageNumber: metadata.Page,
		PerPage:    metadata.PerPage,
		PageCount:  metadata.PageCount,
		TotalCount: metadata.TotalCount,
		Subjects:   make([]*pb.SubjectMetadata, len(subjects)),
	}

	for i, subject := range subjects {
		resp.Subjects[i] = &pb.SubjectMetadata{
			Id:        subject.Id,
			Name:      subject.Name,
			SubjectId: subject.SubjectId,
			Semester:  subject.Semester,
			Section:   subject.Section,
			Year:      subject.Year,
		}
	}

	return &resp, nil
}

func (s *SubjectService) GetSubjectById(ctx context.Context, req *pb.GetSubjectByIdRequest) (*pb.GetSubjectByIdResponse, error) {
	subject, err := s.subjectRepo.FindSubjectById(req.Id)
	if err != nil {
		switch err {
		case entity.ErrNotFound:
			return nil, status.Error(codes.NotFound, "not found")
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	prerequisitesResp := make([]*pb.Subject, 0)
	for _, prerequisite := range subject.Prerequisites {
		prerequisitesResp = append(prerequisitesResp, &pb.Subject{
			Id:        prerequisite.Id,
			SubjectId: prerequisite.SubjectId,
			Name:      prerequisite.Name,
		})
	}

	instructorResp := make([]*pb.Instructor, 0)
	for _, instructor := range subject.Instructors {
		instructorResp = append(instructorResp, &pb.Instructor{
			Id:       instructor.Id,
			FullName: instructor.FullName,
		})
	}

	return &pb.GetSubjectByIdResponse{
		Subject: &pb.Subject{
			Id:            subject.Id,
			SubjectId:     subject.SubjectId,
			Name:          subject.Name,
			Semester:      subject.Semester,
			Section:       subject.Section,
			Year:          subject.Year,
			Faculty:       subject.Faculty,
			Description:   subject.Description,
			Prerequisites: prerequisitesResp,
			Instructors:   instructorResp,
		},
	}, nil
}

func (s *SubjectService) CreateSubject(ctx context.Context, req *pb.CreateSubjectRequest) (*pb.CreateSubjectResponse, error) {
	if req.SubjectId == "" {
		return nil, status.Error(codes.InvalidArgument, "subject ID not provided")
	}
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "subject name not provided")
	}
	if req.Semester < 1 || req.Semester > 3 {
		return nil, status.Error(codes.InvalidArgument, "subject semester should be 1,2 or 3")
	}
	if req.Section < 1 || req.Section > 100 {
		return nil, status.Error(codes.InvalidArgument, "subject section should be between 1 and 100")
	}
	if req.Year < 2000 {
		return nil, status.Error(codes.InvalidArgument, "subject year must be after 1999")
	}
	if req.Faculty == "" {
		return nil, status.Error(codes.InvalidArgument, "subject faculty not provided")
	}

	prerequisites := make([]entity.Subject, 0)
	for _, prerequisiteId := range req.PrerequisiteIds {
		prerequisites = append(prerequisites, entity.Subject{Id: prerequisiteId})
	}

	instructors := make([]entity.Instructor, 0)
	for _, instructorId := range req.InstructorIds {
		instructors = append(instructors, entity.Instructor{Id: instructorId})
	}

	subject := entity.Subject{
		SubjectId:     req.SubjectId,
		Name:          req.Name,
		Semester:      req.Semester,
		Section:       req.Section,
		Year:          req.Year,
		Faculty:       req.Faculty,
		Description:   req.Description,
		Prerequisites: prerequisites,
		Instructors:   instructors,
	}

	if err := s.subjectRepo.CreateSubject(&subject); err != nil {
		fmt.Println(err)
		switch err {
		case entity.ErrConstraintViolation:
			return nil, status.Error(codes.InvalidArgument, "bad request")
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	prerequisitesResp := make([]*pb.Subject, 0)
	for _, prerequisite := range subject.Prerequisites {
		prerequisitesResp = append(prerequisitesResp, &pb.Subject{
			Id:        prerequisite.Id,
			SubjectId: prerequisite.SubjectId,
			Name:      prerequisite.Name,
		})
	}

	instructorResp := make([]*pb.Instructor, 0)
	for _, instructor := range subject.Instructors {
		instructorResp = append(instructorResp, &pb.Instructor{
			Id:       instructor.Id,
			FullName: instructor.FullName,
		})
	}

	return &pb.CreateSubjectResponse{
		Subject: &pb.Subject{
			Id:            subject.Id,
			SubjectId:     subject.SubjectId,
			Name:          subject.Name,
			Semester:      subject.Semester,
			Section:       subject.Section,
			Year:          subject.Year,
			Faculty:       subject.Faculty,
			Description:   subject.Description,
			Prerequisites: prerequisitesResp,
			Instructors:   instructorResp,
		},
	}, nil
}

func (s *SubjectService) UpdateSubject(ctx context.Context, req *pb.UpdateSubjectRequest) (*pb.UpdateSubjectResponse, error) {
	return nil, nil
}

func (s *SubjectService) DeleteSubject(ctx context.Context, req *pb.DeleteSubjectRequest) (*pb.DeleteSubjectResponse, error) {
	return nil, nil
}
