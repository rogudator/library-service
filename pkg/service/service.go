// Это пакет промежуточного слоя между слоем Repository и слоем Handlers.
package service

import "github.com/rogudator/library-service/pkg/repository"

// Внизу специальный комментарий Mockgen для автогенерации моков. Моки понадобятся для написания unit тестов.
//go:generate mockgen -source=service.go -destination=mocks/mock.go

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

// Конструктор слоя service, который по большей части
// вызывает соответствующие методы слоя repository.
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Books:   NewBooksServce(repos.Books),
		Authors: NewAuthorsServce(repos.Authors),
	}
}
