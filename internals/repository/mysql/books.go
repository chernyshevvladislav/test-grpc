package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/chernyshevvladislav/test-grpc/internals/domain"
)

type BooksRepo struct {
	db *sql.DB
}

func NewBooksRepo(db *sql.DB) *BooksRepo {
	return &BooksRepo{db: db}
}

func (rep *BooksRepo) GetBooksByAuthorName(ctx context.Context, name string) ([]domain.Book, error) {
	books := make([]domain.Book, 0)
	if name == "" {
		return books, nil
	}
	name = fmt.Sprintf("%%%s%%", name)
	r, err := rep.db.QueryContext(ctx, "SELECT books.* FROM books INNER JOIN author_book on author_book.book_id = books.id INNER JOIN authors on authors.id = author_book.author_id WHERE CONCAT(authors.last_name, ' ',authors.first_name, ' ',authors.patronymic) LIKE ?", name)

	if err != nil {
		fmt.Errorf("Error when trying get books: %v", err)
		return nil, err
	}
	for r.Next() {
		book := domain.Book{}
		err = r.Scan(&book.ID, &book.Title)
		books = append(books, book)
	}

	return books, nil
}
