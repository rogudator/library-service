package handlers

import (
	"errors"

	"github.com/rogudator/library-service/pkg/service"
	"github.com/rogudator/library-service/proto/libraryServicePb"
)

// Список возможных ошибок:
var (
	ErrNoResults = errors.New("there are no results")
	ErrNoBooks   = errors.New("there are no books by this author")
	ErrNoAuthors = errors.New("book has no author")
)

// В предыдущих слоях (service и repository) в основном была только
// одно строчка с указателем или на repository, или на бд.
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
