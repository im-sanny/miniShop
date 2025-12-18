package util

import (
	"encoding/json"
	"log"
	"net/http"
)

// SendData writes the given data as JSON with the specified status code.
func SendData(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}

// SendError sends a standardized JSON error response.
func SendError(w http.ResponseWriter, statusCode int, message string) {
	SendData(w, statusCode, map[string]string{"error": message})
}
