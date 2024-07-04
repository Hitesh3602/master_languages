package main

import (
	"log"
	"net/http"

	"github.com/Hitesh3602/master_languages/internal/config"
	"github.com/Hitesh3602/master_languages/internal/db"
	"github.com/Hitesh3602/master_languages/internal/service"
	transportHttp "github.com/Hitesh3602/master_languages/internal/transport"
)

func main() {
	cfg := config.LoadConfig()

	// Connect to the database
	database, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer database.Close()

	// Initialize repository and service
	languageRepo := db.NewPostgresLanguageRepository(database)
	languageService := service.NewLanguageService(languageRepo)

	// Set up HTTP transport
	handler := transportHttp.NewHTTPHandler(languageService)

	// Start the HTTP server
	log.Println("Starting server on :8083")
	if err := http.ListenAndServe(":8083", handler); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
