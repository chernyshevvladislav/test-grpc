package v1

import (
	"github.com/chernyshevvladislav/test-grpc/internals/domain"
	"github.com/chernyshevvladislav/test-grpc/pkg/generate/library"
)

func SerializeToGetBooksByAuthorResponse(books []domain.Book) *library.GetBooksByAuthorNameResponse {
	Books := make([]*library.Book, 0, len(books))
	for _, book := range books {
		Books = append(Books, &library.Book{Title: book.Title})
	}
	resp := &library.GetBooksByAuthorNameResponse{
		Books: Books,
	}

	return resp
}

func SerializeToGetAuthorsByBookTitleResponse(authors []domain.Author) *library.GetAuthorsByBookTitleResponse {
	var Authors []*library.Author
	for _, author := range authors {
		authorName := author.LastName + " " + author.FirstName + " " + author.Patronymic
		Authors = append(Authors, &library.Author{Name: authorName})
	}
	resp := &library.GetAuthorsByBookTitleResponse{
		Authors: Authors,
	}

	return resp
}
