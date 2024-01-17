package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	// "gg/pb"

	"net"

	pb "gg/protos"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func ConnectDB() {
	HOST := "localhost"
	PORT := "54321"
	USER := "sm_user3"
	PASSWORD := "12345678"
	DBNAME := "testdb"
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DBNAME,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "Slow SQL Log: ", log.LstdFlags),
		logger.Config{
			SlowThreshold:             100 * time.Millisecond,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  false,
		},
	)

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{Logger: newLogger})

	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&Book{})
	Db = db
}

func GetDB() *gorm.DB {
	if Db == nil {
		ConnectDB()
	}
	return Db
}

type server struct {
	pb.UnimplementedGreeterServer
	pb.UnimplementedBookServer
}
type Book struct {
	ID       int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name"`
	AuthorID int64  `json:"author_id"`
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello " + req.Name}, nil
}

func (s *server) CreateBook(ctx context.Context, req *pb.CreateBookRequestBody) (*pb.CreateBookResponseBody, error) {
	// Implement your logic to create a book
	// For demonstration purposes, just return a dummy response
	return &pb.CreateBookResponseBody{
		Id:  123,
		Msg: "Book created successfully",
	}, nil
}

func (s *server) GetBook(ctx context.Context, req *pb.GetBookRequestBody) (*pb.GetBookResponseBody, error) {
	// Implement your logic to get a book by ID
	// For demonstration purposes, just return a dummy response
	return &pb.GetBookResponseBody{
		Id:       456,
		Name:     "Sample Book",
		AuthorId: 789,
	}, nil
}

func (s *server) UpdateBook(ctx context.Context, req *pb.CreateBookRequestBody) (*pb.CreateBookResponseBody, error) {
	// Implement your logic to update a book
	// For demonstration purposes, just return a dummy response
	return &pb.CreateBookResponseBody{
		Id:  123,
		Msg: "Book updated successfully",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ConnectDB()

	s := grpc.NewServer()
	reflection.Register(s)

	// Register each gRPC service separately
	pb.RegisterGreeterServer(s, &server{})
	pb.RegisterBookServer(s, &server{})

	fmt.Println("gRPC server is running on port 50051")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
