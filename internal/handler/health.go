package handler

import (
	"net/http"

	"github.com/snowchest67/todo-api/internal/model"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := model.HealthResponse{Status: "ok"}
	sendJSON(w, response)
}