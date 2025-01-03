package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/steliosmagalios/link-shortener/internal/database"
	"github.com/steliosmagalios/link-shortener/internal/database/repository"
)

type LinkHandler struct {
	queries *repository.Queries
	db      *database.Database
}

func NewLinkHandler(q *repository.Queries, db *database.Database) *LinkHandler {
	return &LinkHandler{
		queries: q,
		db:      db,
	}
}

func (h *LinkHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.queries.GetAllLinks(h.db.Ctx)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(data)
}
