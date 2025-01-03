package handlers

import (
	"context"
	"net/http"

	"github.com/steliosmagalios/link-shortener/internal/database/repository"
)

type LinkHandler struct {
	queries *repository.Queries
	ctx     context.Context
}

func NewLinkHandler(ctx context.Context, q *repository.Queries) *LinkHandler {
	return &LinkHandler{
		queries: q,
		ctx:     ctx,
	}
}

func (h *LinkHandler) FindAll(w http.ResponseWriter, r *http.Request) {}

func (h *LinkHandler) FindBySlug(w http.ResponseWriter, r *http.Request) {}

func (h *LinkHandler) CreateOne(w http.ResponseWriter, r *http.Request) {}

func (h *LinkHandler) UpdateOne(w http.ResponseWriter, r *http.Request) {}

func (h *LinkHandler) DeleteOne(w http.ResponseWriter, r *http.Request) {}
