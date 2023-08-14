package grpc

import (
	"context"
	"fmt"
	"github.com/chernyshevvladislav/test-grpc/internals/service"
	"github.com/chernyshevvladislav/test-grpc/pkg/generate/library"
)

type Handler struct {
	BooksService   service.Book
	AuthorsService service.Author
	library.UnimplementedLibraryServer
}

func NewHandler(bookService service.Book, authorService service.Author) *Handler {
	return &Handler{
		BooksService:   bookService,
		AuthorsService: authorService,
	}
}

func (h *Handler) GetAuthorsByBookTitle(ctx context.Context, req *library.GetAuthorsByBookTitleRequest) (*library.GetAuthorsByBookTitleResponse, error) {
	resp, err := h.AuthorsService.FindAuthorsByBookTitle(ctx, req.Title)
	if err != nil {
		return nil, fmt.Errorf("cannot find Authors: %w", err)
	}

	return resp, nil
}

func (h *Handler) GetBooksByAuthorName(ctx context.Context, req *library.GetBooksByAuthorNameRequest) (*library.GetBooksByAuthorNameResponse, error) {
	resp, err := h.BooksService.FindBooksByAuthorName(ctx, req.Name)
	if err != nil {
		return nil, fmt.Errorf("cannot find Authors: %w", err)
	}

	return resp, nil
}
