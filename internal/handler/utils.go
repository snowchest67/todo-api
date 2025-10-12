package handler

import (
	"encoding/json"
	"net/http"
)

func sendJSON(w http.ResponseWriter, msg interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(msg); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}