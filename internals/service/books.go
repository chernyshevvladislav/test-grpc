package service

import (
	"context"
	v1 "github.com/chernyshevvladislav/test-grpc/internals/delivery/grpc/v1"
	"github.com/chernyshevvladislav/test-grpc/internals/repository"
	"github.com/chernyshevvladislav/test-grpc/pkg/generate/library"
)

type Book struct {
	repo repository.Books
}

func NewBooksService(repo repository.Books) Book {
	return Book{repo: repo}
}

func (s *Book) FindBooksByAuthorName(ctx context.Context, title string) (*library.GetBooksByAuthorNameResponse, error) {
	books, err := s.repo.GetBooksByAuthorName(ctx, title)
	if err != nil {
		return nil, err
	}
	return v1.SerializeToGetBooksByAuthorResponse(books), nil
}
