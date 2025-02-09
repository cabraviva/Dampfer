package api

import (
	"Dampfer/docker"
	"Dampfer/utils"
	"encoding/json"
	"io"
	"net/http"
)

func ImageInspect(w http.ResponseWriter, r *http.Request, username string) {
	if r.Body == nil {
		utils.Log.Info("Request body is empty")
		http.Error(w, "Empty request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Check and decode JSON
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.Log.Error("Error reading request body: ", err)
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	if !json.Valid(body) {
		utils.Log.Info("Invalid JSON")
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var payload IdPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		utils.Log.Error("Error decoding JSON: ", err)
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Validate fields
	if payload.Id == "" {
		utils.Log.Info("Missing required fields in JSON payload")
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	images, err := docker.ImageInspect(payload.Id)

	if err != nil {
		http.Error(w, "Failed to inspect image", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(images)
}

func ImageDelete(w http.ResponseWriter, r *http.Request, username string) {
	// Get the query parameters
	repository := r.URL.Query().Get("repository")
	tag := r.URL.Query().Get("tag")

	if repository == "" || tag == "" {
		http.Error(w, "Both 'repository' and 'tag' query parameters are required", http.StatusBadRequest)
		return
	}

	// Download the image and store it in the database
	success, err := docker.ImageRm(repository, tag)
	if err != nil {
		utils.Log.Warn("Failed to delete image: " + err.Error())
		json.NewEncoder(w).Encode(false)

		return
	}

	json.NewEncoder(w).Encode(success)
}
