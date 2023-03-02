package service

import "github.com/rogudator/library-service/pkg/repository"

type AuthorsService struct {
	repo repository.Authors
}

func NewAuthorsServce(repo repository.Authors) *AuthorsService {
	return &AuthorsService{
		repo: repo,
	}
}

func (s *AuthorsService) GetAuthorsOfBook(bookName string) ([]string, error) {
	return s.repo.GetAuthorsOfBook(bookName)
}
