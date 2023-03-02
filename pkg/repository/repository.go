package repository

import "github.com/jmoiron/sqlx"

type Book interface {
	GetBookByAuthor(authorName string) ([]string, error)
}

type Author interface {
	GetAuthorOfBook(bookName string) ([]string, error)
}

type Repository struct {
	Book
	Author
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Book:   NewBookMysql(db),
		Author: NewAuthorMysql(db),
	}
}
