package main

import (
	"net"

	"google.golang.org/grpc"

	"github.com/jKulrativid/SA-Subject-Service/src/app/service"
	"github.com/jKulrativid/SA-Subject-Service/src/repository"

	"github.com/jKulrativid/SA-Subject-Service/src/database"
	pb "github.com/jKulrativid/SA-Subject-Service/src/grpc/subject"
)

func main() {
	dbConn, err := database.NewDatabaseConnection()
	if err != nil {
		panic(err)
	}

	if err := database.Migrate(dbConn); err != nil {
		panic(err)
	}

	subjectRepo := repository.NewSubjectRepository(dbConn)
	instructorRepo := repository.NewInstructorRepository(dbConn)

	subjectService := service.NewSubjectService(subjectRepo)
	instructorService := service.NewInstructorService(instructorRepo)

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
