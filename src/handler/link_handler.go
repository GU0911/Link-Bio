package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"linkbio-go/src/model"
	"linkbio-go/src/repository"

	"github.com/gorilla/mux"
)

// LinkHandler handles HTTP requests related to links
type LinkHandler struct {
	Repo *repository.LinkRepository
}

// NewLinkHandler creates a new instance of LinkHandler
func NewLinkHandler(repo *repository.LinkRepository) *LinkHandler {
	return &LinkHandler{Repo: repo}
}

// GetAllLinks retrieves all links
func (h *LinkHandler) GetAllLinks(w http.ResponseWriter, r *http.Request) {
	links, err := h.Repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(links)
}

// CreateLink creates a new link
func (h *LinkHandler) CreateLink(w http.ResponseWriter, r *http.Request) {
	var link model.Link
	if err := json.NewDecoder(r.Body).Decode(&link); err != nil {
		http.Error(w, "Bad Request: Invalid JSON format", http.StatusBadRequest)
		return
	}
	if link.Title == "" || link.URL == "" {
		http.Error(w, "Bad Request: Title and URL cannot be empty", http.StatusBadRequest)
		return
	}
	newLink, err := h.Repo.Create(link)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newLink)
}

// GetLinkByID retrieves a link by its ID from the URL
func (h *LinkHandler) GetLinkByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid link ID", http.StatusBadRequest)
		return
	}

	link, err := h.Repo.GetByID(id)
	if err == sql.ErrNoRows {
		http.Error(w, "Link not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(link)
}

// UpdateLink updates an existing link
func (h *LinkHandler) UpdateLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid link ID", http.StatusBadRequest)
		return
	}

	var link model.Link
	if err := json.NewDecoder(r.Body).Decode(&link); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	if link.Title == "" || link.URL == "" {
		http.Error(w, "Title and URL cannot be empty", http.StatusBadRequest)
		return
	}

	updatedLink, err := h.Repo.Update(id, link)
	if err == sql.ErrNoRows {
		http.Error(w, "Link not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedLink)
}

// DeleteLink deletes a link by its ID
func (h *LinkHandler) DeleteLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid link ID", http.StatusBadRequest)
		return
	}

	err = h.Repo.Delete(id)
	if err == sql.ErrNoRows {
		http.Error(w, "Link not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}