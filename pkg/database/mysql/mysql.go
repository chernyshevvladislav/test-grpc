package mysql

import (
	"context"
	"database/sql"
	"github.com/chernyshevvladislav/test-grpc/internals/config"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Client interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

func NewClient(cfg config.MySQL) (*sql.DB, error) {
	path := cfg.User + ":" + cfg.Password + "@tcp(127.0.0.1:" + cfg.Port + ")/" + cfg.Database
	db, err := sql.Open("mysql", path)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, nil
}
