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

	// Links subdomain
	linkHandler := handlers.NewLinkHandler(queries, &db)
	mux.HandleFunc("/links", linkHandler.GetAll)

	return mux
}
