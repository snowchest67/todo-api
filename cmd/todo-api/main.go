package main

import (
	"encoding/json"
	"net/http"
)

type Response map[string]string

func sendJSON(w http.ResponseWriter, msg interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(msg); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		response := Response{"message": "Welcome to Todo API!"}
		sendJSON(w, response)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request){
		response := Response{"status": "ok"}
		sendJSON(w, response)
	})

	http.ListenAndServe(":8080", nil)
}