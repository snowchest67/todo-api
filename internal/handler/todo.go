package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/snowchest67/todo-api/internal/model"
)

var (
	todos   []model.Todo
	nextID  = 1
)

func TodosHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        getTodos(w, r)
    case http.MethodPost:
        createTodo(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}


func getTodos(w http.ResponseWriter, r *http.Request) {
	sendJSON(w, todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {

	var req model.CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	todo := model.Todo{ID: nextID, Title: req.Title}
	todos = append(todos, todo)
	nextID++

	w.WriteHeader(http.StatusCreated)
	sendJSON(w, todo)
}

func GetTodoByIDHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Incorrect number of elements in the request", http.StatusBadRequest)
		return
	}
	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for _, todo := range todos {
		if todo.ID == id {
			w.WriteHeader(http.StatusOK)
			sendJSON(w, todo)
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound)
}