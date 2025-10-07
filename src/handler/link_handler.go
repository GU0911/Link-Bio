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

type LinkHandler struct {
	Repo     repository.ILinkRepository
	Logger   *slog.Logger
	Validate *validator.Validate
}

func NewLinkHandler(repo repository.ILinkRepository, logger *slog.Logger, validate *validator.Validate) *LinkHandler {
	return &LinkHandler{
		Repo:     repo,
		Logger:   logger,
		Validate: validate,
	}
}

func (h *LinkHandler) writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// ... (semua fungsi handler lainnya tetap sama persis)

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

func (h *LinkHandler) GetAllLinks(w http.ResponseWriter, r *http.Request) {
	links, err := h.Repo.GetAll()
	if err != nil {
		h.Logger.Error("failed to get all links", "error", err)
		h.writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
		return
	}
	h.writeJSON(w, http.StatusOK, links)
}

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