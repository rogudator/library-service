package handlers

import (
	"context"

	"github.com/rogudator/library-service/proto/libraryServicePb"
)

func (rpc *RPC) GetBooksByAuthor(ctx context.Context, req *libraryServicePb.BooksRequest) (*libraryServicePb.BooksResponse, error) {
	authorName := req.GetAuthorName()
	if authorName == "" {
		return nil, ErrWrongName
	}
	books, err := rpc.services.GetBooksByAuthor(authorName)
	if err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, ErrNoBooks
	}
	res := libraryServicePb.BooksResponse{
		Books: books,
	}

	return &res, nil
}
