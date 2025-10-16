package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/snowchest67/todo-api/internal/model"
	"github.com/snowchest67/todo-api/internal/repository"
)

type TodoHandler struct {
    repo repository.TodoRepository
}

func NewTodoHandler(repo *repository.PostgresRepo) *TodoHandler {
    return &TodoHandler{repo: repo}
}

func (h *TodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/todos" {

        switch r.Method {
        case http.MethodGet:
            h.getTodos(w, r)
        case http.MethodPost:
            h.createTodo(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    } else if strings.HasPrefix(r.URL.Path, "/todos/") {
        switch r.Method {
        case http.MethodGet:
            h.getTodoByID(w, r)
        case http.MethodDelete:
            h.deleteTodo(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    } else {
        http.NotFound(w, r)
    }
}


func (h *TodoHandler) getTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.repo.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	sendJSON(w, todos)
}

func (h *TodoHandler) createTodo(w http.ResponseWriter, r *http.Request) {

	var req model.CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	id, err := h.repo.CreateTodo(r.Context(), req.Title)
	if err != nil {
		http.Error(w, "Failed to create todo", http.StatusInternalServerError)
    return
  }

	todo := model.Todo{ID: id, Title: req.Title}
  w.WriteHeader(http.StatusCreated)
  sendJSON(w, todo)
}

func(h *TodoHandler) getTodoByID(w http.ResponseWriter, r *http.Request) {
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

	todo, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	if todo == nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	sendJSON(w, todo)	
}

func(h *TodoHandler) deleteTodo(w http.ResponseWriter, r *http.Request) {
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

	err = h.repo.DeleteByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

