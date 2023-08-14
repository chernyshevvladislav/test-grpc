package repository

import (
	"context"
	"database/sql"
	"github.com/chernyshevvladislav/test-grpc/internals/domain"
	"github.com/chernyshevvladislav/test-grpc/internals/repository/mysql"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type Books interface {
	GetBooksByAuthorName(ctx context.Context, name string) ([]domain.Book, error)
}

type Authors interface {
	GetAuthorsByBookTitle(ctx context.Context, title string) ([]domain.Author, error)
}

type books struct {
	Books
}

type authors struct {
	Authors
}

func NewBooks(db *sql.DB) Books {
	return &books{mysql.NewBooksRepo(db)}
}
func NewAuthors(db *sql.DB) Authors {
	return &authors{mysql.NewAuthorsRepo(db)}
}
