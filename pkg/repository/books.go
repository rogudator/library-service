package repository

import (
	"github.com/jmoiron/sqlx"
)

type BookMysql struct {
	db *sqlx.DB
}

// Конструктор структуры реализующей методы описанные в
// интерфейсе Books из файла repository.go.
func NewBookMysql(db *sqlx.DB) *BookMysql {
	return &BookMysql{db: db}
}

// Эта функция делает запрос в базу данных и выводит книги заданного автора.
func (r *BookMysql) GetBooksByAuthor(authorName string) ([]string, error) {

	query := `
	SELECT b.name  FROM library l 
	INNER JOIN authors a 
	ON l.id_author = a.id 
	INNER JOIN books b 
	ON l.id_book = b.id
	WHERE a.name = ?;
	`
	rows, err := r.db.DB.Query(query, authorName)
	if err != nil {
		return nil, err
	}

	var books []string
	for rows.Next() {
		var name string
		if err = rows.Scan(&name); err != nil {
			return nil, err
		}
		books = append(books, name)
	}

	return books, nil
}
