package handlers

import (
	"context"

	"github.com/rogudator/library-service/proto/libraryServicePb"
)

// При обращении к этому методу gRPC сначала проверяется,
// введено ли имя автора, только потом идет запрос в базу данных.
// После успешного запроса, если у автора нет книг, то выводится ошибка.
// Если все прошло хорошо, будут выведены все книги автора.
func (rpc *RPC) GetBooksByAuthor(ctx context.Context, req *libraryServicePb.BooksRequest) (*libraryServicePb.BooksResponse, error) {
	authorName := req.GetAuthorName()
	if authorName == "" {
		return nil, ErrNoResults
	}
	books, err := rpc.services.GetBooksByAuthor(authorName)
	if err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, ErrAuthorNotPresent
	}
	res := libraryServicePb.BooksResponse{
		Books: books,
	}

	return &res, nil
}
