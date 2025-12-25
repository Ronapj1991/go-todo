package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Todo API is running\n")
}

func main() {
	store := NewTodoStore()

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/todos", TodosHandler(store))
	http.HandleFunc("/todos/", TodosByIDHandler(store))

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
