package main

import (
	"go-rest-api/internal/handlers"
	"go-rest-api/internal/services"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	router := mux.NewRouter()

	nasaService := services.NewNASAService()
	neoHandler := handlers.NewNEOHandler(nasaService)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	router.HandleFunc("/neo", neoHandler.GetNEOs).Methods("GET")

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
