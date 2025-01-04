package api

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/steliosmagalios/link-shortener/internal/api/handlers"
	"github.com/steliosmagalios/link-shortener/internal/database/repository"
)

func New(ctx context.Context, conn *pgx.Conn) http.Handler {
	queries := repository.New(conn) // Initialize queries

	mux := http.NewServeMux()

	// Links subroutes
	linkHandler := handlers.NewLinkHandler(ctx, queries)
	mux.HandleFunc("GET /links", linkHandler.FindAll)
	mux.HandleFunc("POST /links", linkHandler.InsertOne)

	mux.HandleFunc("GET /links/{slug}", linkHandler.FindBySlug)
	mux.HandleFunc("PUT /links/{slug}", linkHandler.UpdateOne)
	mux.HandleFunc("DELETE /links/{slug}", linkHandler.DeleteOne)

	return mux
}
