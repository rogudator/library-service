package handlers

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rogudator/library-service/pkg/handlers"
	"github.com/rogudator/library-service/pkg/service"
	mock_service "github.com/rogudator/library-service/pkg/service/mocks"
	"github.com/rogudator/library-service/proto/libraryServicePb"
	"github.com/stretchr/testify/assert"
)

// В данном тесте GetAuthorsOfBook есть 4 вида поведения, когда:
// 1. У книги несколько авторов
// 2. У книги один автор
// 3. Пустые входные данные
// 4. Входных данных нет в библеотеке
func TestGetAuthorsOfBook(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthors, book string)
	testTable := []struct {
		name         string
		bookName     string
		mockBehavior mockBehavior
		authors      *libraryServicePb.AuthorsResponse
		err          error
	}{
		{
			name:     "Get successfully multiple authors",
			bookName: "Chemistry",
			mockBehavior: func(s *mock_service.MockAuthors, book string) {
				s.EXPECT().GetAuthorsOfBook(book).Return([]string{"Karen C. Timberlake", "William Timberlake"}, nil)
			},
			authors: &libraryServicePb.AuthorsResponse{
				Authors: []string{"Karen C. Timberlake", "William Timberlake"},
			},
			err: nil,
		},
		{
			name:     "Get successfully one author",
			bookName: "Anna Karenina",
			mockBehavior: func(s *mock_service.MockAuthors, book string) {
				s.EXPECT().GetAuthorsOfBook(book).Return([]string{"Leo Tolstoy"}, nil)
			},
			authors: &libraryServicePb.AuthorsResponse{
				Authors: []string{"Leo Tolstoy"},
			},
			err: nil,
		},
		{
			name:         "Empty input",
			bookName:     "",
			mockBehavior: func(s *mock_service.MockAuthors, book string) {},
			authors:      nil,
			err:          handlers.ErrEmptyInput,
		},
		{
			name:     "Wrong input",
			bookName: "Kemistry",
			mockBehavior: func(s *mock_service.MockAuthors, book string) {
				s.EXPECT().GetAuthorsOfBook(book).Return(nil, handlers.ErrAuthorNotPresent)
			},
			authors: nil,
			err:     handlers.ErrAuthorNotPresent,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAuthors(c)
			test.mockBehavior(repo, test.bookName)

			services := &service.Service{Authors: repo}
			rpc := handlers.NewRPC(services)

			output, err := rpc.GetAuthorsOfBook(context.Background(), &libraryServicePb.AuthorsRequest{
				BookName: test.bookName,
			})
			assert.Equal(t, output, test.authors)
			assert.Equal(t, err, test.err)
		})
	}
}
