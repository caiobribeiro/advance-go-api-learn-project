package main

import (
	"fmt"
	"log"
	"net/http"

	dbconfig "github.com/caiobribeiro/advance-go-api-learn-project/internal/db_config"
	"github.com/caiobribeiro/advance-go-api-learn-project/internal/handlers"
	"github.com/caiobribeiro/advance-go-api-learn-project/internal/routes"
	serverconfig "github.com/caiobribeiro/advance-go-api-learn-project/server_config"
)

func main() {

	// Load config
	config, err := serverconfig.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config %v", err)
	}

	// connect to the database
	db := dbconfig.ConnectDb(config.DatabaseURL)
	defer db.Close()

	// Create a new handler
	handler := handlers.NewHandlers()

	// set up the HTTP server
	serverMux := http.NewServeMux()

	// Setup Routes
	routes.SetupRoutes(serverMux, handler)

	serverAddress := fmt.Sprintf(":%s", config.ServerPort)
	server := &http.Server{
		Addr:    serverAddress,
		Handler: serverMux,
	}

	fmt.Printf("🚀 Server is up and running on PORT %s\n", serverAddress)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed %v", err)
	}
}
