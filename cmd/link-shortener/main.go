package main

import (
	"context"
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

	ctx := context.Background()

	// Initialize database
	conn := database.NewDatabase(ctx, os.Getenv("DATABASE_URL"))
	defer conn.Close(ctx)

	// Setup router
	router := http.NewServeMux()
	router.Handle("/api/v1/", http.StripPrefix("/api/v1", api.New(ctx, conn)))
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // for healthcheck
		w.WriteHeader(200)
	})

	// Setup and start server
	srv := server.New(os.Getenv("PORT"), router)
	srv.Start()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error occured while reading environment variables. Error:", err)
	}
}
