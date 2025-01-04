package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5"
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

func (h *LinkHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.queries.FindAll(h.ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (h *LinkHandler) FindBySlug(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if slug == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := h.queries.FindBySlug(h.ctx, slug)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (h *LinkHandler) InsertOne(w http.ResponseWriter, r *http.Request) {
	body := repository.InsertOneParams{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if body.Location == "" || body.Slug == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := h.queries.InsertOne(h.ctx, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (h *LinkHandler) UpdateOne(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if slug == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body := repository.UpdateLocationParams{Slug: slug}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if body.Location == "" || body.Slug == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := h.queries.UpdateLocation(h.ctx, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)

}

func (h *LinkHandler) DeleteOne(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if slug == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := h.queries.DeleteBySlug(h.ctx, slug)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
