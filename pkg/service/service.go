package service

import "github.com/rogudator/library-service/pkg/repository"

type Book interface {
	GetBookByAuthor(authorName string) ([]string, error)
}

type Author interface {
	GetAuthorOfBook(bookName string) ([]string, error)
}

type Service struct {
	Book
	Author
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Book:   NewBookServce(repos.Book),
		Author: NewAuthorServce(repos.Author),
	}
}