package routes

import (
	"Dampfer/docker"
	"encoding/json"
	"net/http"
)

func ListAllContainers(w http.ResponseWriter, r *http.Request) {
	containers, err := docker.ListAllContainers()

	if err != nil {
		http.Error(w, "Failed to list all containers", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(containers)
}
