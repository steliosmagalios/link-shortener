package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	Conn *pgx.Conn
	Ctx  context.Context
}

func NewDatabase(connString string) Database {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connString)

	if err != nil {
		log.Fatalln("Error occured while initializing database. Error: ", err)
	}

	return Database{
		Conn: conn,
		Ctx:  ctx,
	}
}
