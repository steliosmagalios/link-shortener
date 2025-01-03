package main

import (
	"log"
	"net/http"
	"os"

	"github.com/lpernett/godotenv"
	"github.com/steliosmagalios/link-shortener/internal/api"
	"github.com/steliosmagalios/link-shortener/internal/database"
	"github.com/steliosmagalios/link-shortener/internal/server"
)

func main() {
	loadEnv() // Load env variables

	// Initialize database
	db := database.NewDatabase(os.Getenv("DATABASE_URL"))

	// Setup router
	router := http.NewServeMux()
	router.Handle("/api/v1/", http.StripPrefix("/api/v1", api.New(db)))

	// Setup and start server
	srv := server.New(os.Getenv("APP_ADDR"), router)
	srv.Start()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error occured while reading environment variables. Error:", err)
	}
}
