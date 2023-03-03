package handlers

import (
	"errors"

	"github.com/rogudator/library-service/pkg/service"
	"github.com/rogudator/library-service/proto/libraryServicePb"
)

// Список возможных ошибок:
var (
	ErrIncorrectInput   = errors.New("incorrect input")
	ErrAuthorNotPresent = errors.New("library does not have this author")
	ErrBookNotPresent   = errors.New("library does not have this book")
)

// В предыдущих слоях (service и repository) в основном была только
// одна строчка с указателем или на repository, или на бд.
// Здесь помимо указателя на слой service, еще есть интерфейс
// сгенерированный протобафом. В нем описаны методы, которые
// будут реализованы в файлах authors.go и books.go в рамках
// пакета handlers.
type RPC struct {
	services *service.Service
	libraryServicePb.LibraryServiceServer
}

func NewRPC(services *service.Service) *RPC {
	return &RPC{services: services}
}
