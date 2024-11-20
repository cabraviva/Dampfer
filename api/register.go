package api

import (
	"Dampfer/auth"
	"fmt"
	"net/http"
)

// stores all registered endpoints
var RegisteredEndpoints []string

// registerAPI registers an API endpoint with a specific function and allowed method,
// and adds it to the list of endpoints for listing at /api/endpoints.
// Leave minimumPermission empty ("") if no permission is required
func Register(path string, handler func(http.ResponseWriter, *http.Request, string), method string, requiresAuth bool, minimumPermission string) {
	// Record the endpoint in the format "METHOD /path"
	RegisteredEndpoints = append(RegisteredEndpoints, fmt.Sprintf("%s %s", method, path))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		if requiresAuth {
			if isAuthorized, errMsg, username := auth.IsRequestAuthorized(w, r, minimumPermission); isAuthorized {
				w.Header().Set("Content-Type", "application/json")
				handler(w, r, username)
			} else {
				if errMsg == "" {
					errMsg = "Unauthorized"
				}
				http.Error(w, errMsg, http.StatusUnauthorized)
				return
			}
		} else {
			w.Header().Set("Content-Type", "application/json")
			handler(w, r, "")
		}
	})
}
