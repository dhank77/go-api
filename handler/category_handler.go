package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"go-api/models"
	"go-api/service"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(s *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: s}
}

func (h *CategoryHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := strings.TrimPrefix(r.URL.Path, "/categories")
	path = strings.TrimPrefix(path, "/")

	if path == "" {
		switch r.Method {
		case "GET":
			h.GetAll(w, r)
		case "POST":
			h.Create(w, r)
		default:
			http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
		}
	} else {
		id, err := strconv.Atoi(path)
		if err != nil {
			http.Error(w, `{"error":"Invalid ID"}`, http.StatusBadRequest)
			return
		}
		switch r.Method {
		case "GET":
			h.GetByID(w, r, id)
		case "PUT":
			h.Update(w, r, id)
		case "DELETE":
			h.Delete(w, r, id)
		default:
			http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
		}
	}
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetAll()
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categories)
}

func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request, id int) {
	cat, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, `{"error":"Category not found"}`, http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(cat)
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var cat models.Category
	if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
		http.Error(w, `{"error":"Invalid request body"}`, http.StatusBadRequest)
		return
	}
	if err := h.service.Create(&cat); err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cat)
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request, id int) {
	var cat models.Category
	if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
		http.Error(w, `{"error":"Invalid request body"}`, http.StatusBadRequest)
		return
	}
	cat.ID = id
	if err := h.service.Update(&cat); err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cat)
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request, id int) {
	if err := h.service.Delete(id); err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Category deleted"})
}
