package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HelloHandler handles the root endpoint
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from the backend!")
}

// StatusHandler handles the API status endpoint
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
}
