package service

import "github.com/rogudator/library-service/pkg/repository"

type Books interface {
	GetBooksByAuthor(authorName string) ([]string, error)
}

type Authors interface {
	GetAuthorsOfBook(bookName string) ([]string, error)
}

type Service struct {
	Books
	Authors
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Books:   NewBooksServce(repos.Books),
		Authors: NewAuthorsServce(repos.Authors),
	}
}