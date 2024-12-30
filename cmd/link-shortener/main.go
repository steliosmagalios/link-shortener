package main

import (
	"encoding/json"
	"os"

	"github.com/lpernett/godotenv"
	v1 "github.com/steliosmagalios/link-shortener/internal/api/v1"
	"github.com/steliosmagalios/link-shortener/internal/database"
	"github.com/steliosmagalios/link-shortener/internal/database/repository"
	"github.com/steliosmagalios/link-shortener/internal/server"
)

func main() {
	loadEnv()

	db, _ := database.NewDatabase(os.Getenv("DATABASE_URL"))
	queries := repository.New(db.Conn)

	res, _ := queries.GetAllLinks(db.Ctx)
	o, _ := json.MarshalIndent(res, "", "\t")
	println(string(o))

	server := server.NewServer(":8888", v1.NewAPI())
	server.ListenAndServe()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load environment variables!")
	}
}
