package service

import "github.com/rogudator/library-service/pkg/repository"

type AuthorService struct {
	repo repository.Author
}

func NewAuthorServce(repo repository.Author) *AuthorService {
	return &AuthorService{
		repo: repo,
	}
}

func (s *AuthorService) GetAuthorOfBook(bookName string) ([]string, error) {
	return s.repo.GetAuthorOfBook(bookName)
}