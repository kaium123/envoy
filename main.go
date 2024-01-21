package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"net"

	pb "gg/app/protos"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func readConfig() {
	var err error

	viper.SetConfigFile("base.env")
	viper.SetConfigType("props")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	if _, err := os.Stat("base.env"); os.IsNotExist(err) {
		fmt.Println("WARNING: file base.env not found")
	} else {
		viper.SetConfigFile("base.env")
		viper.SetConfigType("props")
		err = viper.MergeInConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	for _, key := range viper.AllKeys() {
		viper.BindEnv(key)
	}

}

func ConnectDB() {
	// HOST := "localhost"
	// PORT := "54321"
	// USER := "sm_user3"
	// PASSWORD := "12345678"
	// DBNAME := "testdb"
	HOST := viper.GetString("HOST")
	PORT := viper.GetString("PORT")
	USER := viper.GetString("DB_USER")
	PASSWORD := viper.GetString("PASSWORD")
	DBNAME := viper.GetString("DBNAME")
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
		log.Fatal(err)
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
	db := GetDB()
	book := &Book{Name: req.Name, AuthorID: int64(req.AuthorId)}
	err := db.Create(book).Error
	if err != nil {
		return nil, err
	}

	return &pb.CreateBookResponseBody{
		Id:  int32(book.ID),
		Msg: "Book created successfully",
	}, nil
}

func (s *server) GetBook(ctx context.Context, req *pb.BookID) (*pb.GetBookResponseBody, error) {
	db := GetDB()
	book := &Book{}
	err := db.Where("id = ? ", req.Id).First(book).Error
	if err != nil {
		return nil, err
	}

	return &pb.GetBookResponseBody{
		Id:       int32(book.ID),
		Name:     book.Name,
		AuthorId: int32(book.AuthorID),
	}, nil
}

func (s *server) UpdateBook(ctx context.Context, req *pb.UpdateBookRequestBody) (*pb.CreateBookResponseBody, error) {
	db := GetDB()
	book := &Book{ID: int64(req.Id), Name: req.Name, AuthorID: int64(req.AuthorId)}
	err := db.Save(book).Error
	if err != nil {
		return nil, err
	}

	return &pb.CreateBookResponseBody{
		Id:  int32(book.ID),
		Msg: "Book updated successfully",
	}, nil
}

func main() {
	readConfig()
	port := viper.GetString("APP_PORT")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ConnectDB()

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterGreeterServer(s, &server{})
	pb.RegisterBookServer(s, &server{})

	fmt.Println("gRPC server is running on port " + port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
