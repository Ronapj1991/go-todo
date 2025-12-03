package main

import (
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
	}
}
