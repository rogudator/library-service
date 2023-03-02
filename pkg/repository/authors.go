package repository

import (
	"github.com/jmoiron/sqlx"
)

type AuthorMysql struct {
	db *sqlx.DB
}

// Конструктор стуктуры реализующей методы описанные в
// интерфейсе Authors из файла repository.go
func NewAuthorMysql(db *sqlx.DB) *AuthorMysql {
	return &AuthorMysql{db: db}
}

// Эта функция делает запрос в базу данных и выводит авторов по названию книги.
func (r *AuthorMysql) GetAuthorsOfBook(bookName string) ([]string, error) {
	var books []string

	query := `
	select a.name  from library l 
	inner join authors a 
	on l.id_author = a.id 
	inner join books b 
	on l.id_book = b.id
	WHERE b.name = ?;
	`
	rows, err := r.db.DB.Query(query, bookName)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var name string
		if err = rows.Scan(&name); err != nil {
			return nil, err
		}
		books = append(books, name)
	}

	return books, nil
}
