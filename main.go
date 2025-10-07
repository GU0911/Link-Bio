package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"linkbio-go/src/config"
	"linkbio-go/src/handler"
	"linkbio-go/src/middleware"
	"linkbio-go/src/repository"
	"linkbio-go/src/util"

	"github.com/gorilla/mux"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	cfg := config.LoadConfig()
	validate := util.NewValidator()

	db := config.ConnectDB(cfg.DB)
	defer db.Close()

	linkRepo := repository.NewLinkRepository(db)
	linkHandler := handler.NewLinkHandler(linkRepo, logger, validate)

	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	api.Use(middleware.LoggingMiddleware(logger))

	api.HandleFunc("/links", linkHandler.GetAllLinks).Methods("GET")
	api.HandleFunc("/links", linkHandler.CreateLink).Methods("POST")
	api.HandleFunc("/links/{id:[0-9]+}", linkHandler.GetLinkByID).Methods("GET")
	api.HandleFunc("/links/{id:[0-9]+}", linkHandler.UpdateLink).Methods("PUT")
	api.HandleFunc("/links/{id:[0-9]+}", linkHandler.DeleteLink).Methods("DELETE")

	serverAddr := ":" + cfg.ServerPort
	logger.Info("server starting", "address", serverAddr)

	err := http.ListenAndServe(serverAddr, r)
	if err != nil {
		logger.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}