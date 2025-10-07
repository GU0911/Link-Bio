package main

import (
	"fmt"
	"log"
	"net/http"

	"linkbio-go/src/config"
	"linkbio-go/src/handler"
	"linkbio-go/src/repository"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Failed to load .env file. Using system environment variables.")
	}

	db := config.ConnectDB()
	defer db.Close()

	linkRepo := repository.NewLinkRepository(db)
	linkHandler := handler.NewLinkHandler(linkRepo)

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Welcome to the Link Bio API with Go!"}`))
	}).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/links", linkHandler.GetAllLinks).Methods("GET")
	api.HandleFunc("/links", linkHandler.CreateLink).Methods("POST")
	api.HandleFunc("/links/{id:[0-9]+}", linkHandler.GetLinkByID).Methods("GET")
	api.HandleFunc("/links/{id:[0-9]+}", linkHandler.UpdateLink).Methods("PUT")
	api.HandleFunc("/links/{id:[0-9]+}", linkHandler.DeleteLink).Methods("DELETE")

	port := "8080"
	fmt.Printf("Server is running at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}