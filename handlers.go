package main

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

type CreateTodoRequest struct {
	Description string `json:"description"`
}

func parseIDFromPath(path string) (int, error) {
	urlToString := strings.Split(path, "/")
	var cleanPath []string

	for _, v := range urlToString {
		if v != "" {
			cleanPath = append(cleanPath, v)
		}
	}

	if len(cleanPath) == 0 {
		return 0, errors.New("Invalid path")
	}

	ID, err := strconv.Atoi(cleanPath[len(cleanPath)-1])
	if err != nil {
		return 0, errors.New("No ID found")
	}
	return ID, nil
}

func UpdateTodoHandler(store *TodoStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPatch {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		id, err := parseIDFromPath(r.URL.Path)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var updates map[string]interface{}
		if err = json.NewDecoder(r.Body).Decode(&updates); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if err = store.UpdateTodoByID(id, updates); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		updatedTodo, ok := store.FindTodoByID(id)
		if !ok {
			http.Error(w, "Todo not found", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedTodo)
	}
}

func CreateTodoHandler(store *TodoStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
			return
		}

		var req CreateTodoRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if strings.TrimSpace(req.Description) == "" {
			http.Error(w, "Input cannot be empty", http.StatusBadRequest)
			return
		}

		newTodo := store.AddTodo(req.Description)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newTodo)
	}
}

func ListTodoHandler(store *TodoStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		todos := store.GetTodos()

		w.Header().set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todos)
	}
}
