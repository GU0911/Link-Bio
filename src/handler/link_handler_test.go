package handler

import (
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"linkbio-go/src/model"
	"linkbio-go/src/repository"

	"github.com/gorilla/mux"
)

type mockLinkRepository struct {
	links []model.Link
	err   error
}

func (m *mockLinkRepository) GetAll() ([]model.Link, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.links, nil
}
func (m *mockLinkRepository) GetByID(id int) (model.Link, error) { return model.Link{}, nil }
func (m *mockLinkRepository) Create(link model.Link) (model.Link, error) { return model.Link{}, nil }
func (m *mockLinkRepository) Update(id int, link model.Link) (model.Link, error) { return model.Link{}, nil }
func (m *mockLinkRepository) Delete(id int) error { return nil }

func TestGetAllLinks_Success(t *testing.T) {
	// Arrange
	mockRepo := &mockLinkRepository{
		links: []model.Link{
			{ID: 1, Title: "Test", URL: "http://test.com", CreatedAt: time.Now()},
		},
		err: nil,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	handler := NewLinkHandler(mockRepo, logger, nil) // validator is not needed for this test

	req, err := http.NewRequest("GET", "/api/links", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	
	// Act
	router := mux.NewRouter()
	router.HandleFunc("/api/links", handler.GetAllLinks)
	router.ServeHTTP(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedBody := `[{"id":1,"title":"Test","url":"http://test.com"` // We only check the start
	if !contains(rr.Body.String(), expectedBody) {
		t.Errorf("handler returned unexpected body: got %v", rr.Body.String())
	}
}

func TestGetAllLinks_Failure(t *testing.T) {
	// Arrange
	mockRepo := &mockLinkRepository{
		err: errors.New("database error"),
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	handler := NewLinkHandler(mockRepo, logger, nil)

	req, err := http.NewRequest("GET", "/api/links", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Act
	router := mux.NewRouter()
	router.HandleFunc("/api/links", handler.GetAllLinks)
	router.ServeHTTP(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}

func contains(s, substr string) bool {
    return len(s) >= len(substr) && s[:len(substr)] == substr
}