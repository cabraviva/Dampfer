package auth

import (
	"Dampfer/utils"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// CustomClaims extends jwt.RegisteredClaims to include the permission field
type CustomClaims struct {
	jwt.RegisteredClaims
	Permission string `json:"permission"`
}

// GenerateTokenForUser generates a JWT token for the given username
func GenerateTokenForUser(username, password string) (string, error) {
	// Check if the user exists in the database
	users, err := ListUsers()
	if err != nil {
		utils.Log.Error("Failed to list users: ", err)
		return "", err
	}

	var userExists bool
	var userPermission string
	for _, user := range users {
		if user["username"] == username {
			userExists = true
			userPermission = user["permission"]
			break
		}
	}

	if !userExists {
		utils.Log.Error("While generating token: User does not exist: ", username)
		return "", fmt.Errorf("user %s does not exist", username)
	}

	// Verify the password
	isValidPassword, err := VerifyPassword(username, password)
	if err != nil {
		utils.Log.Error("While generating token: Error verifying password for user: ", username)
		return "", err
	}
	if !isValidPassword {
		utils.Log.Error("While generating token: Invalid password for user: ", username)
		return "", fmt.Errorf("invalid password for user %s", username)
	}

	// Retrieve the secret key from the database
	secretKey, err := GetJWTSecret()
	if err != nil {
		return "", err
	}

	// Create the custom claims with permission
	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Dampfer",                                               // Your app's name or domain
			Subject:   username,                                                // The username as the subject
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)), // Set expiration (30 days)
			IssuedAt:  jwt.NewNumericDate(time.Now()),                          // Automatically sets the issued time if not set
		},
		Permission: userPermission, // Add the permission to the claims
	}

	// Create the token with the custom claims and sign it using HMAC
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		utils.Log.Error("Failed to sign the token: ", err)
		return "", err
	}

	return tokenString, nil
}

// Checks whether a token is valid and also returns username and permissions.
// There will never be an error if return value is true
func ValidateToken(tokenString string) (bool, string, string, error) {
	tokenSignatureValid, username, permissions, err1 := verifyToken(tokenString)

	if err1 != nil {
		// Error while verifying token
		return false, username, permissions, err1
	}

	if !tokenSignatureValid {
		// Token invalid
		return false, username, permissions, nil
	}

	// Check for existing user with same permissions
	claimsValid, err2 := verifyUserExistsAndPermissions(username, permissions)

	if err2 != nil {
		return false, username, permissions, err2
	}

	// Return whether token is valid all-together
	return claimsValid, username, permissions, nil
}

// verifyToken verifies the token and returns the username and permission if valid
func verifyToken(tokenString string) (bool, string, string, error) {
	// Retrieve the secret key from the database
	secretKey, err := GetJWTSecret()
	if err != nil {
		utils.Log.Error("Failed to retrieve secret key: ", err)
		return false, "", "", err
	}

	// Parse the token and validate the claims
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			utils.Log.Error("Unexpected signing method: ", token.Method)
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		utils.Log.Error("Failed to parse token: ", err)
		return false, "", "", err
	}

	// Extract the claims from the token
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		utils.Log.Error("Invalid token claims")
		return false, "", "", fmt.Errorf("invalid token")
	}

	// Return the username and permission from the claims
	return true, claims.Subject, claims.Permission, nil
}

// verifyUserExistsAndPermissions checks if the user exists and has the correct permissions
func verifyUserExistsAndPermissions(username, expectedPermission string) (bool, error) {
	// Query the database for the user
	users, err := ListUsers()
	if err != nil {
		utils.Log.Error("Failed to list users: ", err)
		return false, err
	}

	for _, user := range users {
		if user["username"] == username {
			// Check if the permission matches
			if user["permission"] == expectedPermission {
				return true, nil // User exists and permission matches
			}
			return false, fmt.Errorf("user %s permission mismatch", username)
		}
	}

	// User not found
	utils.Log.Error("User not found: ", username)
	return false, fmt.Errorf("user %s does not exist", username)
}
