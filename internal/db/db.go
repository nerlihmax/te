package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func NewDB() *pgx.Conn {
	db, err := pgx.Connect(context.TODO(), os.Getenv("DB_DSN"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to connect to database %v\n", err)
		os.Exit(1)
	}

	return db
}
