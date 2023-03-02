package handlers

import (
	"errors"

	"github.com/rogudator/library-service/pkg/service"
	"github.com/rogudator/library-service/proto/libraryServicePb"
)

var (
	ErrNoNameEnteredError = errors.New("no name entered")
	ErrNoBooks = errors.New("there are no books by this author")
	ErrNoAuthors = errors.New("book has no author")
)

type RPC struct {
	services *service.Service
	libraryServicePb.LibraryServiceServer
}

func NewRPC(services *service.Service) *RPC {
	return &RPC{services: services}
}