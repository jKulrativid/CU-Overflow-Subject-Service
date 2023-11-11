package main

import (
	"log"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"

	"github.com/jKulrativid/SA-Subject-Service/src/app/service"
	"github.com/jKulrativid/SA-Subject-Service/src/repository"
	"github.com/joho/godotenv"

	"github.com/jKulrativid/SA-Subject-Service/src/database"
	pb "github.com/jKulrativid/SA-Subject-Service/src/grpc/subject"
)

var Godogs int

func main() {
	err := godotenv.Load()

	dbPort, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		log.Fatal("Error converting DB port to an integer")
	}

	dbConfig := database.DatabaseConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     dbPort,
		SslMode:  os.Getenv("POSTGRES_SSL_MODE"),
	}

	dbConn, err := database.NewDatabaseConnection(&dbConfig)
	if err != nil {
		panic(err)
	}

	if err := database.Migrate(dbConn); err != nil {
		panic(err)
	}

	subjectRepo := repository.NewSubjectRepository(dbConn, 10)
	instructorRepo := repository.NewInstructorRepository(dbConn, 10)

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
