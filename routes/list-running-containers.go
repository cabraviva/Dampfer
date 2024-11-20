package routes

import (
	"Dampfer/docker"
	"encoding/json"
	"net/http"
)

func ListRunningContainers(w http.ResponseWriter, r *http.Request, username string) {
	containers, err := docker.ListRunningContainers()

	if err != nil {
		http.Error(w, "Failed to list running containers", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(containers)
}
