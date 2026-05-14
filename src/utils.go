package src

import (
	"encoding/json"
	"net/http"
)

// writeJSON converts a Go value to JSON and sends it to the client
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// ReadJSON reads JSON from the request body and converts it to a Go value
func ReadJSON(r *http.Request, dest any) error {
	return json.NewDecoder(r.Body).Decode(dest)
}