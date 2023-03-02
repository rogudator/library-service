package service

import "github.com/rogudator/library-service/pkg/repository"

type BookService struct {
	repo repository.Book
}

func NewBookServce(repo repository.Book) *BookService {
	return &BookService{
		repo: repo,
	}
}

func (s *BookService) GetBookByAuthor(authorName string) ([]string, error) {
	return s.repo.GetBookByAuthor(authorName)
}