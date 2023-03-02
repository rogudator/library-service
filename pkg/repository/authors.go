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
	
	err := checkBook(r, bookName); if err != nil {
		return nil, err
	}

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

// Проверка, есть ли книга в библиотеке.
func checkBook(r *AuthorMysql, bookName string) error{
	query := `
	SELECT name FROM books
	WHERE name = ?;
	`

	rows, err := r.db.DB.Query(query, bookName)
	if err != nil {
		return err
	}
	
	var name string
	for rows.Next() {
	if err = rows.Scan(&name); err != nil {
		return err
	}
}
	if name == "" {
		return ErrBookNotInLibrary
	}

	return nil
}