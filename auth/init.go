package auth

import "Dampfer/database"

var db = database.DB

func init() {
	// Check for or generate a new secret for jwt
	initializeSecretKey()
}
