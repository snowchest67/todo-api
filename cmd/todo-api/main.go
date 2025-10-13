package main

import (
	"net/http"

	"github.com/snowchest67/todo-api/internal/handler"
)


func main() {

	http.HandleFunc("/", handler.RootHandler)

	http.HandleFunc("/health", handler.HealthHandler)

	http.HandleFunc("/todos", handler.TodosHandler)

	http.HandleFunc("/todos/", handler.TodoByIDHandler)

	http.ListenAndServe(":8080", nil)
}