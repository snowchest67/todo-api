package handler

import (
	"net/http"

	"github.com/snowchest67/todo-api/internal/model"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	response := model.RootResponse{Message: "Welcome to Todo API!"}
	sendJSON(w, response)
}