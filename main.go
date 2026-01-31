package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var categories = []Category{}
var nextID = 1

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/categories", categoriesHandler)
	http.HandleFunc("/categories/", categoryByIDHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func categoriesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(categories)

	case "POST":
		var cat Category
		json.NewDecoder(r.Body).Decode(&cat)
		cat.ID = nextID
		nextID++
		categories = append(categories, cat)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(cat)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func categoryByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	index := -1
	for i, cat := range categories {
		if cat.ID == id {
			index = i
			break
		}
	}

	switch r.Method {
	case "GET":
		if index == -1 {
			http.Error(w, "Category not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(categories[index])

	case "PUT":
		if index == -1 {
			http.Error(w, "Category not found", http.StatusNotFound)
			return
		}
		var cat Category
		json.NewDecoder(r.Body).Decode(&cat)
		cat.ID = id
		categories[index] = cat
		json.NewEncoder(w).Encode(cat)

	case "DELETE":
		if index == -1 {
			http.Error(w, "Category not found", http.StatusNotFound)
			return
		}
		categories = append(categories[:index], categories[index+1:]...)
		json.NewEncoder(w).Encode(map[string]string{"message": "Category deleted"})

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
