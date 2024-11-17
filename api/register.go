package api

import (
	"fmt"
	"net/http"
)

// stores all registered endpoints
var RegisteredEndpoints []string

// registerAPI registers an API endpoint with a specific function and allowed method,
// and adds it to the list of endpoints for listing at /api/endpoints.
func Register(path string, handler func(http.ResponseWriter, *http.Request), method string) {
	// Record the endpoint in the format "METHOD /path"
	RegisteredEndpoints = append(RegisteredEndpoints, fmt.Sprintf("%s %s", method, path))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		handler(w, r)
	})
}
