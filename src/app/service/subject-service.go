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

func SubjectToPb(subject *entity.Subject) *pb.Subject {
	sectionIds := make([]int64, 0)
	for _, section := range subject.Sections {
		sectionIds = append(sectionIds, section.Id)
	}

	return &pb.Subject{
		Id:            subject.Id,
		SubjectId:     subject.SubjectId,
		Name:          subject.Name,
		Semester:      subject.Semester,
		SectionIds:    sectionIds,
		Year:          subject.Year,
		Faculty:       subject.Faculty,
		Description:   subject.Description,
		Prerequisites: subject.Prerequisites,
	}
}

func SectionToPb(section *entity.Section) *pb.Section {
	instructorIds := make([]int64, 0)
	for _, instructor := range section.Instructors {
		instructorIds = append(instructorIds, instructor.Id)
	}

	return &pb.Section{
		Id:            section.Id,
		SubjectId:     section.SubjectId,
		Number:        section.Number,
		Description:   section.Description,
		InstructorIds: instructorIds,
	}
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
	if req.YearRangeStart != 0 && req.YearRangeStop != 0 {
		isValidYearRange := req.YearRangeStart <= req.YearRangeStop
		if isValidYearRange {
			query["year_range_start"] = req.YearRangeStart
			query["year_range_stop"] = req.YearRangeStop
		}
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
		Subjects:   make([]*pb.SubjectMetadata, 0),
	}

	for _, subject := range subjects {
		resp.Subjects = append(resp.Subjects, &pb.SubjectMetadata{
			Id:        subject.Id,
			Name:      subject.Name,
			SubjectId: subject.SubjectId,
			Semester:  subject.Semester,
			Year:      subject.Year,
		})
	}

	return &resp, nil
}

func (s *SubjectService) ValidateSubjectId(ctx context.Context, req *pb.ValidateSubjectIdRequest) (*pb.ValidateSubjectIdResponse, error) {
	_, err := s.subjectRepo.FindSubjectById(req.Id)
	if err != nil {
		switch err {
		case entity.ErrNotFound:
			return &pb.ValidateSubjectIdResponse{Valid: false}, nil
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &pb.ValidateSubjectIdResponse{Valid: true}, nil
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

	return &pb.GetSubjectByIdResponse{Subject: SubjectToPb(subject)}, nil
}

func (s *SubjectService) CreateSubject(ctx context.Context, req *pb.CreateSubjectRequest) (*pb.CreateSubjectResponse, error) {
	if !req.IsAdmin {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}
	if req.SubjectId == "" {
		return nil, status.Error(codes.InvalidArgument, "subject ID not provided")
	}

	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "subject name not provided")
	}
	if req.Semester < 1 || req.Semester > 3 {
		return nil, status.Error(codes.InvalidArgument, "subject semester should be 1,2 or 3")
	}
	if req.Year < 2000 {
		return nil, status.Error(codes.InvalidArgument, "subject year must be after 1999")
	}
	if req.Faculty == "" {
		return nil, status.Error(codes.InvalidArgument, "subject faculty not provided")
	}

	subject := entity.Subject{
		SubjectId:   req.SubjectId,
		Name:        req.Name,
		Semester:    req.Semester,
		Year:        req.Year,
		Faculty:     req.Faculty,
		Description: req.Description,
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

	return &pb.CreateSubjectResponse{Subject: SubjectToPb(&subject)}, nil
}

func (s *SubjectService) UpdateSubject(ctx context.Context, req *pb.UpdateSubjectRequest) (*pb.UpdateSubjectResponse, error) {
	if !req.IsAdmin {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}

	if req.Semester != 0 && (req.Semester < 1 || req.Semester > 3) {
		return nil, status.Error(codes.InvalidArgument, "subject semester should be 1,2 or 3")
	}
	if req.Year != 0 && (req.Year < 2000) {
		return nil, status.Error(codes.InvalidArgument, "subject year must be after 1999")
	}

	subject := entity.Subject{
		Id:            req.Id,
		SubjectId:     req.SubjectId,
		Name:          req.Name,
		Semester:      req.Semester,
		Year:          req.Year,
		Faculty:       req.Faculty,
		Description:   req.Description,
		Prerequisites: req.Prerequisites,
	}

	err := s.subjectRepo.UpdateSubject(&subject)
	if err != nil {
		switch err {
		case entity.ErrNotFound:
			return nil, status.Error(codes.NotFound, "subject with given ID not found")
		case entity.ErrConstraintViolation:
			return nil, status.Error(codes.InvalidArgument, "bad request")
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &pb.UpdateSubjectResponse{Subject: SubjectToPb(&subject)}, nil
}

func (s *SubjectService) DeleteSubject(ctx context.Context, req *pb.DeleteSubjectRequest) (*pb.DeleteSubjectResponse, error) {
	if !req.IsAdmin {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}

	if req.Id < 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid ID")
	}

	subject, err := s.subjectRepo.DeleteSubjectById(req.Id)
	if err != nil {
		switch err {
		case entity.ErrNotFound:
			return nil, status.Error(codes.NotFound, "subject with given ID not found")
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &pb.DeleteSubjectResponse{Subject: SubjectToPb(subject)}, nil
}

func (s *SubjectService) ValidateSection(ctx context.Context, req *pb.ValidateSectionRequest) (*pb.ValidateSectionResponse, error) {
	_, err := s.subjectRepo.GetSectionByNumberAndSubjectId(req.SectionNumber, req.SubjectId, req.Year, req.Semester)
	if err != nil {
		switch err {
		case entity.ErrNotFound:
			return &pb.ValidateSectionResponse{Valid: false}, nil
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &pb.ValidateSectionResponse{Valid: true}, nil
}

func (s *SubjectService) CreateSection(ctx context.Context, req *pb.CreateSectionRequest) (*pb.CreateSectionResponse, error) {
	if !req.IsAdmin {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}

	if req.SubjectId == 0 {
		return nil, status.Error(codes.InvalidArgument, "subject id must be provided")
	}
	if req.Number < 1 || req.Number > 100 {
		return nil, status.Error(codes.InvalidArgument, "section number must be 1-100")
	}

	section := entity.Section{
		SubjectId:   req.SubjectId,
		Number:      req.Number,
		Description: req.Description,
	}

	err := s.subjectRepo.CreateSection(&section)
	if err != nil {
		switch err {
		case entity.ErrNotFound:
			return nil, status.Error(codes.InvalidArgument, "subject with given subject id not found")
		case entity.ErrConstraintViolation:
			return nil, status.Error(codes.InvalidArgument, "subject already has section with given number")
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &pb.CreateSectionResponse{Section: SectionToPb(&section)}, nil
}

func (s *SubjectService) UpdateSection(ctx context.Context, req *pb.UpdateSectionRequest) (*pb.UpdateSectionResponse, error) {
	if !req.IsAdmin {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}

	if req.Number < 1 || req.Number > 100 {
		return nil, status.Error(codes.InvalidArgument, "section number must be 1-100")
	}

	section := entity.Section{
		Id:          req.Id,
		Number:      req.Number,
		Description: req.Description,
	}

	err := s.subjectRepo.UpdateSection(&section)
	if err != nil {
		switch err {
		case entity.ErrNotFound:
			return nil, status.Error(codes.InvalidArgument, "subject with given subject id not found")
		case entity.ErrConstraintViolation:
			return nil, status.Error(codes.InvalidArgument, "subject already has section with given number")
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &pb.UpdateSectionResponse{Section: SectionToPb(&section)}, nil
}

func (s *SubjectService) DeleteSection(ctx context.Context, req *pb.DeleteSectionRequest) (*pb.DeleteSectionResponse, error) {
	if !req.IsAdmin {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}

	if req.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "section ID not provided")
	}

	section, err := s.subjectRepo.DeleteSection(req.Id)
	if err != nil {
		switch err {
		case entity.ErrNotFound:
			return nil, status.Error(codes.InvalidArgument, "section with given ID not found")
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &pb.DeleteSectionResponse{Section: SectionToPb(section)}, nil
}
