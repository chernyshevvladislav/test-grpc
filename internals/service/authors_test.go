package service

import (
	"context"
	"github.com/chernyshevvladislav/test-grpc/internals/domain"
	mock_service "github.com/chernyshevvladislav/test-grpc/internals/repository/mocks"
	"github.com/chernyshevvladislav/test-grpc/pkg/generate/library"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestFindAuthorsByBookTitle(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthors, title string)

	testTable := []struct {
		title        string
		mockBehavior mockBehavior
		expected     *library.GetAuthorsByBookTitleResponse
	}{
		{
			title: "Преступление и наказание",
			mockBehavior: func(s *mock_service.MockAuthors, title string) {
				s.EXPECT().GetAuthorsByBookTitle(context.TODO(), title).Return([]domain.Author{
					{
						ID:         1,
						FirstName:  "Федор",
						LastName:   "Достоевский",
						Patronymic: "Михайлович",
					},
				}, nil)
			},
			expected: &library.GetAuthorsByBookTitleResponse{
				Authors: []*library.Author{{Name: "Достоевский Федор Михайлович"}},
			},
		},
	}

	for _, test := range testTable {
		t.Run(test.title, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			repo := mock_service.NewMockAuthors(c)
			test.mockBehavior(repo, test.title)
			service := &Author{repo: repo}
			authors, err := service.FindAuthorsByBookTitle(context.TODO(), test.title)
			if err != nil {
				t.Fail()
			}

			for i := 0; i < len(test.expected.Authors); i++ {
				expected := test.expected.Authors[i].Name
				actual := authors.Authors[i].Name

				if expected != actual {
					t.Fail()
				}
			}
		})
	}
}
