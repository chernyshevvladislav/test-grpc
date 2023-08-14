package app

import (
	"github.com/chernyshevvladislav/test-grpc/internals/config"
	"github.com/chernyshevvladislav/test-grpc/internals/delivery/grpc"
	"github.com/chernyshevvladislav/test-grpc/internals/repository"
	"github.com/chernyshevvladislav/test-grpc/internals/server"
	"github.com/chernyshevvladislav/test-grpc/internals/service"
	"github.com/chernyshevvladislav/test-grpc/pkg/database/mysql"
	"log"
)

func Run() {
	cfg := config.New()
	db, err := mysql.NewClient(cfg.MySQL)
	if err != nil {
		log.Fatalf("Cannot start server: %s", err)
	}
	booksService := service.NewBooksService(repository.NewBooks(db))
	authorsService := service.NewAuthorsService(repository.NewAuthors(db))
	handler := grpc.NewHandler(booksService, authorsService)
	s := server.NewServer(cfg.HTTP, handler)
	err = s.Start()
	if err != nil {
		log.Fatalf("Cannot start server: %s", err)
	}
}
