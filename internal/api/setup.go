package api

import (
	"net/http"

	"github.com/steliosmagalios/link-shortener/internal/api/handlers"
	"github.com/steliosmagalios/link-shortener/internal/database"
	"github.com/steliosmagalios/link-shortener/internal/database/repository"
)

func New(db database.Database) http.Handler {
	queries := repository.New(db.Conn) // Initialize queries

	mux := http.NewServeMux()

	// Links subroutes
	linkHandler := handlers.NewLinkHandler(queries, &db)
	mux.HandleFunc("GET /links", linkHandler.FindAll)
	mux.HandleFunc("POST /links", linkHandler.CreateOne)

	mux.HandleFunc("GET /links/{slug}", linkHandler.FindBySlug)
	mux.HandleFunc("PUT /links/{slug}", linkHandler.UpdateOne)
	mux.HandleFunc("DELETE /links/{slug}", linkHandler.DeleteOne)

	return mux
}
