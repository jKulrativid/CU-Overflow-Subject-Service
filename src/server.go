package main

import (
	"net"

	"google.golang.org/grpc"

	"github.com/jKulrativid/SA-Subject-Service/src/app/service"
	pb "github.com/jKulrativid/SA-Subject-Service/src/grpc/subject"
)

func main() {
	subjectService := service.NewSubjectService()
	instructorService := service.NewInstructorService()

	grpcServer := grpc.NewServer()
	pb.RegisterSubjectServiceServer(grpcServer, subjectService)
	pb.RegisterInstructorServiceServer(grpcServer, instructorService)

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic("cannot open TCP connection at 0.0.0.0:8080")
	}

	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
