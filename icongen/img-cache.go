package icongen

import (
	"Dampfer/database"
	"fmt"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"strings"
)

// SaveImage stores an image in the database with a given ID
// SaveImage stores an image in the database with its MIME type
func SaveImage(id string, imageData []byte, mimeType string) error {
	query := `INSERT INTO icon_cache (id, image, mime_type) VALUES (?, ?, ?)`
	_, err := database.DB.Exec(query, id, imageData, mimeType)
	return err
}

// GetImage retrieves an image by ID
func GetImage(id string) ([]byte, string, error) {
	var imageData []byte
	var mimeType string
	query := `SELECT image, mime_type FROM icon_cache WHERE id = ?`
	err := database.DB.QueryRow(query, id).Scan(&imageData, &mimeType)
	if err != nil {
		return nil, "", err
	}
	return imageData, mimeType, nil
}

// DownloadImageToDB downloads an image from an external URL and stores it in the icon_cache table
func DownloadImageToDB(id, imageURL string) error {
	// Download the image
	resp, err := http.Get(imageURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the image content
	imageData, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Get MIME type from Content-Type header
	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		return fmt.Errorf("failed to detect MIME type")
	}

	// If MIME type is a generic type, try to detect from file extension
	if !strings.Contains(contentType, "/") {
		ext := filepath.Ext(imageURL)
		if ext != "" {
			contentType = mime.TypeByExtension(ext)
		}
	}

	// Save the image to the database
	return SaveImage(id, imageData, contentType)
}
