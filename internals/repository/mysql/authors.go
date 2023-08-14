package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/chernyshevvladislav/test-grpc/internals/domain"
)

type AuthorsRepo struct {
	db *sql.DB
}

func NewAuthorsRepo(db *sql.DB) AuthorsRepo {
	return AuthorsRepo{db: db}
}
func (a AuthorsRepo) GetAuthorsByBookTitle(ctx context.Context, title string) ([]domain.Author, error) {
	authors := make([]domain.Author, 0)

	if title == "" {
		return authors, nil
	}

	title = fmt.Sprintf("%%%s%%", title)
	r, err := a.db.QueryContext(ctx, "SELECT authors.* FROM authors INNER JOIN author_book on authors.id = author_book.author_id INNER JOIN books on author_book.book_id = books.id WHERE title LIKE ?", title)
	if err != nil {
		return nil, fmt.Errorf("Error when try get authors: %w", err)
	}
	for r.Next() {
		author := domain.Author{}
		err = r.Scan(&author.ID, &author.FirstName, &author.LastName, &author.Patronymic)
		authors = append(authors, author)
	}

	return authors, nil
}
