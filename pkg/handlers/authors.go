package handlers

import (
	"context"

	"github.com/rogudator/library-service/proto/libraryServicePb"
)

// При обращении к этому методу gRPC сначала проверяется,
// введено ли название книги, только потом идет запрос в базу данных.
// После успешного запроса, если у книги нет авторов, то выводится ошибка.
// Если все прошло хорошо, будут выведены авторы книги.
func (rpc *RPC) GetAuthorsOfBook(ctx context.Context, req *libraryServicePb.AuthorsRequest) (*libraryServicePb.AuthorsResponse, error) {
	bookName := req.GetBookName()
	if bookName == "" {
		return nil, ErrNoResults
	}
	authors, err := rpc.services.GetAuthorsOfBook(bookName)
	if err != nil {
		return nil, err
	}
	if len(authors) == 0 {
		return nil, ErrBookNotPresent
	}
	res := libraryServicePb.AuthorsResponse{
		Authors: authors,
	}

	return &res, nil
}
