package service

import (
	"context"
	v1 "github.com/chernyshevvladislav/test-grpc/internals/delivery/grpc/v1"
	"github.com/chernyshevvladislav/test-grpc/internals/repository"
	"github.com/chernyshevvladislav/test-grpc/pkg/generate/library"
)

type Author struct {
	repo repository.Authors
}

func NewAuthorsService(repo repository.Authors) Author {
	return Author{repo: repo}
}

func (s *Author) FindAuthorsByBookTitle(ctx context.Context, title string) (*library.GetAuthorsByBookTitleResponse, error) {
	authors, err := s.repo.GetAuthorsByBookTitle(ctx, title)
	if err != nil {
		return nil, err
	}
	return v1.SerializeToGetAuthorsByBookTitleResponse(authors), nil
}
