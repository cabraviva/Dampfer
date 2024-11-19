package auth

import (
	"Dampfer/utils"
	"crypto/rand"
	"fmt"
)

// Generate a secure random secret key (32 bytes is a good size)
func generateSecretKey() ([]byte, error) {
	// Generate a random key of 32 bytes
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, fmt.Errorf("failed to generate random secret key: %v", err)
	}
	return key, nil
}

// Initialize the secret key in the database (or retrieve it if it already exists)
func initializeSecretKey() {
	// Check if the secret key is already in the database
	var secretKey []byte
	err := db.QueryRow("SELECT secret_key FROM auth LIMIT 1").Scan(&secretKey)
	if err != nil && err.Error() != "sql: no rows in result set" {
		utils.Log.Panic("Failed to retrieve secret key: ", err)
		panic(err)
	}

	if len(secretKey) == 0 { // If no secret key exists, generate and store it
		secretKey, err = generateSecretKey()
		if err != nil {
			utils.Log.Panic("Failed to generate secret key: ", err)
			panic(err)
		}

		// Store the generated secret key in the database
		_, err = db.Exec("INSERT INTO auth (secret_key) VALUES (?)", secretKey)
		if err != nil {
			utils.Log.Panic("Failed to store secret key in database: ", err)
			panic(err)
		}
		utils.Log.Info("Generated and stored a new secret key.")
	} else {
		utils.Log.Info("Secret key found in database.")
	}
}

// GetJWTSecret retrieves the JWT secret key from the database
func GetJWTSecret() ([]byte, error) {
	// Query the database to get the secret key
	var secretKey []byte
	err := db.QueryRow("SELECT secret_key FROM auth LIMIT 1").Scan(&secretKey)
	if err != nil {
		// Log the error if there's an issue retrieving the secret
		utils.Log.Error("Failed to retrieve secret key from the database: ", err)
		return nil, err
	}

	// Return the retrieved secret key
	utils.Log.Info("Successfully retrieved secret key from the database.")
	return secretKey, nil
}
