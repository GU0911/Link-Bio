package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"linkbio-go/src/model"
	"linkbio-go/src/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// LinkHandler handles HTTP requests for links
type LinkHandler struct {
	Repo     *repository.LinkRepository
	Logger   *slog.Logger
	Validate *validator.Validate
}

// NewLinkHandler creates a new instance of LinkHandler
func NewLinkHandler(repo *repository.LinkRepository, logger *slog.Logger, validate *validator.Validate) *LinkHandler {
	return &LinkHandler{
		Repo:     repo,
		Logger:   logger,
		Validate: validate,
	}
}

// Helper to write JSON response
func (h *LinkHandler) writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// CreateLink creates a new link
func (h *LinkHandler) CreateLink(w http.ResponseWriter, r *http.Request) {
	var link model.Link
	if err := json.NewDecoder(r.Body).Decode(&link); err != nil {
		h.writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	if err := h.Validate.Struct(link); err != nil {
		h.writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	newLink, err := h.Repo.Create(link)
	if err != nil {
		h.Logger.Error("failed to create link", "error", err)
		h.writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
		return
	}

	h.writeJSON(w, http.StatusCreated, newLink)
}

// GetAllLinks retrieves all links
func (h *LinkHandler) GetAllLinks(w http.ResponseWriter, r *http.Request) {
	links, err := h.Repo.GetAll()
	if err != nil {
		h.Logger.Error("failed to get all links", "error", err)
		h.writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
		return
	}
	h.writeJSON(w, http.StatusOK, links)
}

// GetLinkByID retrieves a single link by its ID
func (h *LinkHandler) GetLinkByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	link, err := h.Repo.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.writeJSON(w, http.StatusNotFound, map[string]string{"error": "Link not found"})
		} else {
			h.Logger.Error("failed to get link by ID", "id", id, "error", err)
			h.writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
		}
		return
	}

	h.writeJSON(w, http.StatusOK, link)
}

// UpdateLink updates an existing link
func (h *LinkHandler) UpdateLink(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var link model.Link
	if err := json.NewDecoder(r.Body).Decode(&link); err != nil {
		h.writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	if err := h.Validate.Struct(link); err != nil {
		h.writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	updatedLink, err := h.Repo.Update(id, link)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.writeJSON(w, http.StatusNotFound, map[string]string{"error": "Link not found"})
		} else {
			h.Logger.Error("failed to update link", "id", id, "error", err)
			h.writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
		}
		return
	}
	h.writeJSON(w, http.StatusOK, updatedLink)
}

// DeleteLink deletes a link by its ID
func (h *LinkHandler) DeleteLink(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	err := h.Repo.Delete(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.writeJSON(w, http.StatusNotFound, map[string]string{"error": "Link not found"})
		} else {
			h.Logger.Error("failed to delete link", "id", id, "error", err)
			h.writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
		}
		return
	}
	h.writeJSON(w, http.StatusNoContent, nil)
}