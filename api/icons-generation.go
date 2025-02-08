package api

import (
	"Dampfer/icongen"
	"Dampfer/utils"
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

// GetIcon handles fetching an image by ID
func GetIcon(w http.ResponseWriter, r *http.Request, username string) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	imageData, mimeType, err := icongen.GetImage(id)
	if err != nil {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}

	// Set content type from database
	w.Header().Set("Content-Type", mimeType)
	w.Write(imageData)
}

// UploadIcon handles the image upload
// UploadIcon handles the image upload
func UploadIcon(w http.ResponseWriter, r *http.Request, username string) {
	err := r.ParseMultipartForm(10 << 20) // 10MB max file size
	if err != nil {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("icon")
	if err != nil {
		http.Error(w, "Invalid file upload", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read file content into byte slice
	fileBytes := make([]byte, fileHeader.Size)
	_, err = file.Read(fileBytes)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	// Detect MIME type
	mimeType := fileHeader.Header.Get("Content-Type")
	if mimeType == "" {
		http.Error(w, "Failed to detect MIME type", http.StatusBadRequest)
		return
	}

	// Get image ID from form
	id := r.FormValue("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	// Save image to DB
	err = icongen.SaveImage(id, fileBytes, mimeType)
	if err != nil {
		utils.Log.Error(err)
		http.Error(w, "Failed to save image", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Image uploaded successfully"))
}
