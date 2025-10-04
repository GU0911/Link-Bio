package handler

import (
	"encoding/json"
	"net/http"

	"linkbio-go/src/model"
	"linkbio-go/src/repository"
)

type LinkHandler struct {
	Repo *repository.LinkRepository
}

func NewLinkHandler(repo *repository.LinkRepository) *LinkHandler {
	return &LinkHandler{Repo: repo}
}

func (h *LinkHandler) GetAllLinks(w http.ResponseWriter, r *http.Request) {
	links, err := h.Repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(links)
}

func (h *LinkHandler) CreateLink(w http.ResponseWriter, r *http.Request) {
	var link model.Link
	if err := json.NewDecoder(r.Body).Decode(&link); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if link.Title == "" || link.URL == "" {
		http.Error(w, "Title and URL cannot be empty", http.StatusBadRequest)
		return
	}

	newLink, err := h.Repo.Create(link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newLink)
}