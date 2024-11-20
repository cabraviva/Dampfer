package routes

import (
	"Dampfer/auth"
	"Dampfer/utils"

	"encoding/json"
	"io"
	"net/http"
)

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request, _ string) {
	if r.Method != http.MethodPost {
		utils.Log.Info("Request is not POST")
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

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

	var payload LoginPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		utils.Log.Error("Error decoding JSON: ", err)
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Validate fields
	if payload.Username == "" || payload.Password == "" {
		utils.Log.Info("Missing required fields in JSON payload")
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// payload is json now
	jwtoken, err := auth.GenerateTokenForUser(payload.Username, payload.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	// Send token
	json.NewEncoder(w).Encode(jwtoken)
	return
}
