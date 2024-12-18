package api

import (
	"net/http"
	"strings"
)

// API_listEndpoints lists all registered API endpoints in plain text
func ListEndpoints(w http.ResponseWriter, r *http.Request, username string) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(strings.Join(RegisteredEndpoints, "\n")))
}
