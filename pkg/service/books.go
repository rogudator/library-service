package service

import "github.com/rogudator/library-service/pkg/repository"

type BooksService struct {
	repo repository.Books
}

func NewBooksServce(repo repository.Books) *BooksService {
	return &BooksService{
		repo: repo,
	}
}

func (s *BooksService) GetBooksByAuthor(authorName string) ([]string, error) {
	return s.repo.GetBooksByAuthor(authorName)
}
