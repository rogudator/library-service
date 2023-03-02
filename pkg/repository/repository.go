package repository

import "github.com/jmoiron/sqlx"

type Books interface {
	GetBooksByAuthor(authorName string) ([]string, error)
}

type Authors interface {
	GetAuthorsOfBook(bookName string) ([]string, error)
}

type Repository struct {
	Books
	Authors
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Books:   NewBookMysql(db),
		Authors: NewAuthorMysql(db),
	}
}
