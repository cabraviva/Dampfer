package api

import (
	"net/http"
)

// API_listEndpoints lists all registered API endpoints in plain text
func JWTValidCheckEndpoint(w http.ResponseWriter, r *http.Request, username string) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("SEEMS LIKE EVERYTHING IS FINE."))
}
