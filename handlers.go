package main

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

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
		err = json.NewDecoder(r.Body).Decode(&updates)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		err = store.UpdateTodoByID(id, updates)
		if err != nil {
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
