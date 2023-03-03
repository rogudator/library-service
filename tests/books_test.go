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

// В данном тесте GetBooksByAuthor есть 4 вида поведения, когда:
// 1. У автора несколько книг
// 2. У автора одна книга
// 3. Пустые входные данные
// 4. Входных данных нет в библеотеке
func TestGetBooksByAuthor(t *testing.T) {
	type mockBehavior func(s *mock_service.MockBooks, author string)
	testTable := []struct {
		name         string
		authorName   string
		mockBehavior mockBehavior
		books        *libraryServicePb.BooksResponse
		err          error
	}{
		{
			name:       "Get succesfully multipe books",
			authorName: "William Timberlake",
			mockBehavior: func(s *mock_service.MockBooks, author string) {
				s.EXPECT().GetBooksByAuthor(author).Return([]string{"Chemistry", "Advanced Chemistry"}, nil)
			},
			books: &libraryServicePb.BooksResponse{
				Books: []string{"Chemistry", "Advanced Chemistry"},
			},
			err: nil,
		},
		{
			name:       "Get succesfully on book",
			authorName: "Leo Tolstoy",
			mockBehavior: func(s *mock_service.MockBooks, author string) {
				s.EXPECT().GetBooksByAuthor(author).Return([]string{"Anna Karenina"}, nil)
			},
			books: &libraryServicePb.BooksResponse{
				Books: []string{"Anna Karenina"},
			},
			err: nil,
		},
		{
			name:         "Empty",
			authorName:   "",
			mockBehavior: func(s *mock_service.MockBooks, author string) {},
			books:        nil,
			err:          handlers.ErrNoResults,
		},
		{
			name:       "WrongInput",
			authorName: "William Gajillion",
			mockBehavior: func(s *mock_service.MockBooks, author string) {
				s.EXPECT().GetBooksByAuthor(author).Return(nil, handlers.ErrAuthorNotPresent)
			},
			books: nil,
			err:   handlers.ErrAuthorNotPresent,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockBooks(c)
			test.mockBehavior(repo, test.authorName)

			services := &service.Service{Books: repo}
			rpc := handlers.NewRPC(services)

			output, err := rpc.GetBooksByAuthor(context.Background(), &libraryServicePb.BooksRequest{
				AuthorName: test.authorName,
			})
			assert.Equal(t, output, test.books)
			assert.Equal(t, err, test.err)
		})
	}
}
