package routes

import (
	"Dampfer/docker"
	"encoding/json"
	"net/http"
)

func ListVolumes(w http.ResponseWriter, r *http.Request) {
	volumes, err := docker.ListVolumes()

	if err != nil {
		http.Error(w, "Failed to list volumes", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(volumes)
}
