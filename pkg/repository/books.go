package repository

import (
	"github.com/jmoiron/sqlx"
)

type BookMysql struct {
	db *sqlx.DB
}

func NewBookMysql(db *sqlx.DB) *BookMysql {
	return &BookMysql{db: db}
}

func (r *BookMysql) GetBooksByAuthor(authorName string) ([]string, error) {
	var books []string

	query := `
	select b.name  from library l 
	inner join authors a 
	on l.id_author = a.id 
	inner join books b 
	on l.id_book = b.id
	WHERE a.name = ?;
	`
	rows, err := r.db.DB.Query(query, authorName)
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
