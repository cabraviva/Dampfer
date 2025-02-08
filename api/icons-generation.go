package api

import (
	"Dampfer/icongen"
	"encoding/json"
	"net/http"
)

func IconSearch(w http.ResponseWriter, r *http.Request, username string) {
	// Retrieve the query parameter "q"
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	// Call icongen.GetImagesFor with the query parameter
	links, err := icongen.GetImagesFor(query)
	if err != nil {
		http.Error(w, "Failed to search icons", http.StatusInternalServerError)
		return
	}

	// Encode the result as JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(links)
}
