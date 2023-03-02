// Этот пакет является слоем взаимодействия с базой данных MySQL.
package repository

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

var (
	ErrAuthorNotInLibrary = errors.New("this author is not presented by library")
	ErrBookNotInLibrary = errors.New("this book does not exist in library")
)

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

// Конструктор repository с логикой взаимодействия с бд.
// Потом repository нужен будет для конструктора слоя service.
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Books:   NewBookMysql(db),
		Authors: NewAuthorMysql(db),
	}
}
