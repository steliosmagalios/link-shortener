package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	Conn *pgx.Conn
	Ctx  context.Context
}

func NewDatabase(connString string) (Database, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connString)

	return Database{
		Conn: conn,
		Ctx:  ctx,
	}, err
}
