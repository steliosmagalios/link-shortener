package handlers

import (
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

func (h *LinkHandler) FindAll(w http.ResponseWriter, r *http.Request) {}

func (h *LinkHandler) FindBySlug(w http.ResponseWriter, r *http.Request) {}

func (h *LinkHandler) CreateOne(w http.ResponseWriter, r *http.Request) {}

func (h *LinkHandler) UpdateOne(w http.ResponseWriter, r *http.Request) {}

func (h *LinkHandler) DeleteOne(w http.ResponseWriter, r *http.Request) {}
