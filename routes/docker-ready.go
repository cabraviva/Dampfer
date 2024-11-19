package routes

import (
	"Dampfer/docker"
	"encoding/json"
	"net/http"
)

func DockerReady(w http.ResponseWriter, r *http.Request) {
	response := docker.IsReady()
	json.NewEncoder(w).Encode(response)
}
