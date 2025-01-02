package api

import (
	"net/http"

	"github.com/steliosmagalios/link-shortener/internal/database"
	"github.com/steliosmagalios/link-shortener/internal/database/repository"
	"github.com/steliosmagalios/link-shortener/internal/services"
)

func NewAPI(db *database.Database) http.Handler {
	queries := repository.New(db.Conn)

	mux := http.NewServeMux()

	linkHandler := NewLinkHandler(services.NewLinkService(queries, &db.Ctx))

	mux.HandleFunc("/links", linkHandler.GetAllLinks)

	return mux
}
