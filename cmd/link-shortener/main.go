package main

import (
	"encoding/json"
	"os"

	"github.com/lpernett/godotenv"
	"github.com/steliosmagalios/link-shortener/internal/database"
	"github.com/steliosmagalios/link-shortener/internal/database/repository"
)

func main() {
	loadEnv()

	db, _ := database.NewDatabase(os.Getenv("DATABASE_URL"))
	queries := repository.New(db.Conn)

	res, _ := queries.GetAllLinks(db.Ctx)
	o, _ := json.MarshalIndent(res, "", "\t")
	println(string(o))

}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load environment variables!")
	}
}
