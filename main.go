package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"linkbio-go/src/config"
	"linkbio-go/src/handler"
	"linkbio-go/src/repository"
	"linkbio-go/src/util"

	"github.com/gorilla/mux"
)

func main() {
	// 1. Initialize structured logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// 2. Load configuration
	cfg := config.LoadConfig()

	// 3. Initialize validator
	validate := util.NewValidator()

	// 4. Connect to the database
	db := config.ConnectDB(cfg.DB)
	defer db.Close()

	// 5. Initialize layers
	linkRepo := repository.NewLinkRepository(db)
	linkHandler := handler.NewLinkHandler(linkRepo, logger, validate)

	// 6. Setup router and routes
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/links", linkHandler.GetAllLinks).Methods("GET")
	api.HandleFunc("/links", linkHandler.CreateLink).Methods("POST")
	api.HandleFunc("/links/{id:[0-9]+}", linkHandler.GetLinkByID).Methods("GET")
	api.HandleFunc("/links/{id:[0-9]+}", linkHandler.UpdateLink).Methods("PUT")
	api.HandleFunc("/links/{id:[0-9]+}", linkHandler.DeleteLink).Methods("DELETE")

	// 7. Start the server
	serverAddr := ":" + cfg.ServerPort
	logger.Info("server starting", "address", serverAddr)

	err := http.ListenAndServe(serverAddr, r)
	if err != nil {
		logger.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}