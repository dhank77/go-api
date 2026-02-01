package main

import (
	"fmt"
	"log"
	"net/http"

	"go-api/config"
	"go-api/database"
	"go-api/providers"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Connect to database
	if err := database.Connect(cfg); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()
	fmt.Println("Connected to database!")

	// Register all services
	handlers := providers.RegisterServices()

	// Routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World - Go API with Layered Architecture")
	})
	http.HandleFunc("/categories", handlers.CategoryHandler.Handle)
	http.HandleFunc("/categories/", handlers.CategoryHandler.Handle)
	http.HandleFunc("/products", handlers.ProductHandler.Handle)
	http.HandleFunc("/products/", handlers.ProductHandler.Handle)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
