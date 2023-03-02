package handlers

import (
	"context"

	"github.com/rogudator/library-service/proto/libraryServicePb"
)

func (rpc *RPC) GetAuthorsOfBook(ctx context.Context, req *libraryServicePb.AuthorsRequest) (*libraryServicePb.AuthorsResponse, error) {
	bookName := req.GetBookName()
	if bookName == "" {
		return nil, ErrNoNameEnteredError
	}
	authors, err := rpc.services.GetAuthorsOfBook(bookName)
	if err != nil {
		return nil, err
	}
	if len(authors) == 0 {
		return nil, ErrNoAuthors
	}
	res := libraryServicePb.AuthorsResponse{
		Authors: authors,
	}

	return &res, nil
}