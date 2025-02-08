package icongen

import (
	"Dampfer/database"
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
