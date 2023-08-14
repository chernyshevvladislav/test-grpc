package service

import (
	"context"
	"github.com/chernyshevvladislav/test-grpc/internals/domain"
	mock_service "github.com/chernyshevvladislav/test-grpc/internals/repository/mocks"
	"github.com/chernyshevvladislav/test-grpc/pkg/generate/library"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestFindBooksByAuthorName(t *testing.T) {
	type mockBehavior func(s *mock_service.MockBooks, title string)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		expected     *library.GetBooksByAuthorResponse
	}{
		{
			name: "Достоевский",
			mockBehavior: func(s *mock_service.MockBooks, name string) {
				s.EXPECT().GetBooksByAuthorName(context.TODO(), name).Return([]domain.Book{
					{
						ID:    1,
						Title: "Преступление и наказание",
					},
				}, nil)
			},
			expected: &library.GetBooksByAuthorResponse{
				Books: []*library.Book{{Title: "Преступление и наказание"}},
			},
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			repo := mock_service.NewMockBooks(c)
			test.mockBehavior(repo, test.name)
			service := &Book{repo: repo}
			books, err := service.FindBooksByAuthorName(context.TODO(), test.name)
			if err != nil {
				t.Fail()
			}

			for i := 0; i < len(test.expected.Books); i++ {
				expected := test.expected.Books[i].Title
				actual := books.Books[i].Title
				if expected != actual {
					t.Fail()
				}
			}
		})
	}
}
