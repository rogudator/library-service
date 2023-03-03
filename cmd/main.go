package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/rogudator/library-service/pkg/handlers"
	"github.com/rogudator/library-service/pkg/repository"
	"github.com/rogudator/library-service/pkg/service"
	"github.com/rogudator/library-service/proto/libraryServicePb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting library service...")

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	// Подключение к базе данных.
	db, err := repository.NewMysqlDB(repositoryConfig())
	if err != nil {
		log.Fatal(err)
	}

	// Данный проект был спроектирован в рамках чистой архитектуры.
	// У нас есть три слоя:
	// 1. Repository для связи с базой данных
	repos := repository.NewRepository(db)
	// 2. Services соединяет слой работы с базой данных со слоем работы с gRPC
	services := service.NewService(repos)
	// 3. Handlers нужен для коммуникации по gRPC
	rpc := handlers.NewRPC(services)

	// Здесь мы анонсируем по какому адресу можно обращаться к нашему сервису
	lis, err := net.Listen("tcp", hostname())
	if err != nil {
		log.Fatalf("Failed to listen to: %v", err)
	}

	// Здесь мы создаем пустой grpc сервер в котором зарегистрируем наш gRPC сервис
	s := grpc.NewServer()
	libraryServicePb.RegisterLibraryServiceServer(s, rpc)

	log.Println("Library Service started.")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

// Эта функция возвращает переменную, содержащую заполненный конфиг для подключения к бд
func repositoryConfig() repository.Config {
	if err := godotenv.Load("db.env"); err != nil {
		log.Fatalf("error loading env variavles: %s", err.Error())
	}
	return repository.Config{
		User:     viper.GetString("db.user"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Database: viper.GetString("db.database"),
	}
}

// Получение адреса для сервиса
func hostname() string {
	hostname := fmt.Sprintf("%s:%s", viper.GetString("host"), viper.GetString("port"))
	return hostname
}
