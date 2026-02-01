package main

import (
	"fmt"
	"log"
	"net/http"

	"go-api/database"
	"go-api/providers"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	// Connect to database
	if err := database.InitDB(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

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
