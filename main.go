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
	_ "github.com/spf13/viper/remote"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type ServerConfig struct {
	Address string `mapstructure:"address"`
	Port    string `mapstructure:"port"`
}

type AppConfig struct {
	Database DatabaseConfig `mapstructure:"database"`
	Server   ServerConfig   `mapstructure:"server"`
}

var config AppConfig

func readConfig() {
	consulPath := os.Getenv("CONSUL_PATH")
	consulURL := os.Getenv("CONSUL_URL")

	viper.AddRemoteProvider("consul", consulURL, consulPath)
	viper.SetConfigType("json") // Need to explicitly set this to json

	err := viper.ReadRemoteConfig()
	if err != nil {
		panic(err)
	}

	// Unmarshal the configuration into the AppConfig struct
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Error unmarshalling config:", err)
	}

	viper.AddRemoteProvider("consul", "localhost:8500", "MY_CONSUL_KEY")

	fmt.Printf("Database Host: %s\n", config.Database.Host)
	fmt.Printf("Server Address: %s\n", config.Server.Address)

}

func ConnectDB() {
	// HOST := "localhost"
	// PORT := "54321"
	// USER := "sm_user3"
	// PASSWORD := "12345678"
	// DBNAME := "testdb"
	// HOST := viper.GetString("HOST")
	// PORT := viper.GetString("PORT")
	// USER := viper.GetString("DB_USER")
	// PASSWORD := viper.GetString("PASSWORD")
	// DBNAME := viper.GetString("DBNAME")
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.Username, config.Database.Password, config.Database.DBName,
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
	if !req.IsValid {
		return &pb.CreateBookResponseBody{
			Id:  0,
			Msg: "You are not authorized",
		}, nil
	}
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
	if !req.IsValid {
		return &pb.GetBookResponseBody{
			Id:  0,
			Msg: "You are not authorized",
		}, nil
	}
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
	if !req.IsValid {
		return &pb.CreateBookResponseBody{
			Id:  0,
			Msg: "You are not authorized",
		}, nil
	}
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
	// port := viper.GetString("APP_PORT")
	lis, err := net.Listen("tcp", ":"+config.Server.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ConnectDB()

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterGreeterServer(s, &server{})
	pb.RegisterBookServer(s, &server{})

	fmt.Println("gRPC server is running on port " + config.Server.Port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
