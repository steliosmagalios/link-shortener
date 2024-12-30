package main

import (
	"encoding/json"

	db "github.com/steliosmagalios/link-shortener/internal/database"
	"github.com/steliosmagalios/link-shortener/internal/database/repository"
)

func main() {
	db, _ := db.NewDatabase("postgres://user:password@127.0.0.1:5432/db")
	queries := repository.New(db.Conn)

	res, _ := queries.GetAllLinks(db.Ctx)
	o, _ := json.MarshalIndent(res, "", "\t")
	println(string(o))

}
