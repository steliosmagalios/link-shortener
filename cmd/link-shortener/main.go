package main

import (
	"log"
	"net/http"
	"os"

	"github.com/lpernett/godotenv"
	v1 "github.com/steliosmagalios/link-shortener/internal/api"
	"github.com/steliosmagalios/link-shortener/internal/database"
	"github.com/steliosmagalios/link-shortener/internal/server"
)

func main() {
	loadEnv()

	db, _ := database.NewDatabase(os.Getenv("DATABASE_URL"))

	router := http.NewServeMux()
	router.Handle("/api/v1/", http.StripPrefix("/api/v1", v1.NewAPI(&db)))

	server := server.NewServer(os.Getenv("APP_ADDR"), router)
	log.Fatalln(server.ListenAndServe())
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load environment variables!")
	}
}
