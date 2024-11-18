package routes

import (
	"Dampfer/docker"
	"encoding/json"
	"net/http"
)

func ListImages(w http.ResponseWriter, r *http.Request) {
	images, err := docker.ListImages()

	if err != nil {
		http.Error(w, "Failed to list images", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(images)
}
