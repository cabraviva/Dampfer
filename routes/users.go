package routes

import (
	"Dampfer/auth"
	"Dampfer/utils"
	"encoding/json"
	"io"
	"net/http"
)

type DeleteUserPayload struct {
	Username string `json:"username"`
}

type CreateUserPayload struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Permission string `json:"permission"`
}

type SetPermissionUserPayload struct {
	Username   string `json:"username"`
	Permission string `json:"permission"`
}

type SetPasswordUserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChangePasswordPayload struct {
	Password string `json:"password"`
}

func ListUsers(w http.ResponseWriter, r *http.Request, username string) {
	users, err := auth.ListUsers()

	if err != nil {
		http.Error(w, "Failed to list users", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, username string) {
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

	var payload DeleteUserPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		utils.Log.Error("Error decoding JSON: ", err)
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Validate fields
	if payload.Username == "" {
		utils.Log.Info("Missing required fields in JSON payload: username")
		http.Error(w, "Missing required fields: username", http.StatusBadRequest)
		return
	}

	del_err := auth.DeleteUser(payload.Username)

	if del_err != nil {
		http.Error(w, "Could not delete user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(true)
}

func CreateUser(w http.ResponseWriter, r *http.Request, username string) {
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

	var payload CreateUserPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		utils.Log.Error("Error decoding JSON: ", err)
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Validate fields
	if payload.Username == "" || payload.Password == "" || (payload.Permission == "") {
		utils.Log.Info("Missing required fields in JSON payload")
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	create_err := auth.CreateUser(payload.Username, payload.Password, payload.Permission)

	if create_err != nil {
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(true)
}

func SetUserPermission(w http.ResponseWriter, r *http.Request, username string) {
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

	var payload SetPermissionUserPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		utils.Log.Error("Error decoding JSON: ", err)
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Validate fields
	if payload.Username == "" || (payload.Permission == "") {
		utils.Log.Info("Missing required fields in JSON payload")
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	create_err := auth.SetPermission(payload.Username, payload.Permission)

	if create_err != nil {
		http.Error(w, "Could not change permission", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(true)
}

func SetUserPassword(w http.ResponseWriter, r *http.Request, username string) {
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

	var payload SetPasswordUserPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		utils.Log.Error("Error decoding JSON: ", err)
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Validate fields
	if payload.Username == "" || (payload.Password == "") {
		utils.Log.Info("Missing required fields in JSON payload")
		utils.Log.Info("Received: username=" + payload.Username)
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	create_err := auth.ChangePassword(payload.Username, payload.Password, false) // false because pw changed by admin

	if create_err != nil {
		http.Error(w, "Could not change password", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(true)
}

func ChangeMyPassword(w http.ResponseWriter, r *http.Request, username string) {
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

	var payload ChangePasswordPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		utils.Log.Error("Error decoding JSON: ", err)
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Validate fields
	if payload.Password == "" {
		utils.Log.Info("Missing required fields in JSON payload")
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	create_err := auth.ChangePassword(username, payload.Password, true) // True because pw change should increment

	if create_err != nil {
		http.Error(w, "Could not change password", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(true)
}
