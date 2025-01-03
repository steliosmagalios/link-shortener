package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func NewDatabase(ctx context.Context, connString string) *pgx.Conn {
	conn, err := pgx.Connect(ctx, connString)

	if err != nil {
		log.Fatalln("Error occured while initializing database. Error: ", err)
	}

	return conn
}
