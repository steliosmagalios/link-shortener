package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/steliosmagalios/link-shortener/internal/services"
)

type LinkHandler struct {
	service services.LinkService
}

func NewLinkHandler(service services.LinkService) *LinkHandler {
	return &LinkHandler{
		service: service,
	}
}

func (h *LinkHandler) GetAllLinks(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetAll()
	if err != nil {
		http.Error(w, "An error occured", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (h *LinkHandler) GetLinkFromSlug(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func (h *LinkHandler) CreateLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func (h *LinkHandler) UpdateLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func (h *LinkHandler) DeleteLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
