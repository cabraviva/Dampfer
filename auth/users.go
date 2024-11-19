package auth

import (
	"Dampfer/database"
	"Dampfer/utils"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var db = database.DB

// Permission types
const (
	SystemAdmin = "system-admin"
	Admin       = "admin"
	Insight     = "insight"
)

// hashPassword hashes a plaintext password using bcrypt
func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		utils.Log.Error("Failed to hash password: ", err)
		return "", err
	}
	return string(hashed), nil
}

// checkPassword compares a plaintext password with a hashed password
func checkPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		utils.Log.Error("Password comparison failed: ", err)
	}
	return err == nil
}

// CreateUser adds a new user with a hashed password and permission to the database
func CreateUser(username, password, permission string) error {
	if !isValidPermission(permission) {
		err := errors.New("invalid permission")
		utils.Log.Error("CreateUser failed: ", err)
		return err
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}

	_, err = db.Exec(`INSERT INTO users (username, password_hash, permission) VALUES (?, ?, ?)`,
		username, hashedPassword, permission)
	if err != nil {
		utils.Log.Error("Failed to create user: ", err)
		return err
	}

	utils.Log.Info("User created successfully: ", username)
	return nil
}

// DeleteUser removes a user from the database
func DeleteUser(username string) error {
	_, err := db.Exec(`DELETE FROM users WHERE username = ?`, username)
	if err != nil {
		utils.Log.Error("Failed to delete user: ", err)
		return err
	}

	utils.Log.Info("User deleted successfully: ", username)
	return nil
}

// VerifyPassword checks if the given password matches the stored hash for the user
func VerifyPassword(username, password string) (bool, error) {
	var hashedPassword string
	err := db.QueryRow(`SELECT password_hash FROM users WHERE username = ?`, username).Scan(&hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.Log.Warn("User not found: ", username)
			return false, errors.New("user not found")
		}
		utils.Log.Error("Failed to verify password: ", err)
		return false, err
	}

	return checkPassword(hashedPassword, password), nil
}

// ChangePassword updates the password for an existing user
func ChangePassword(username, newPassword string) error {
	hashedPassword, err := hashPassword(newPassword)
	if err != nil {
		return err
	}

	_, err = db.Exec(`UPDATE users SET password_hash = ? WHERE username = ?`, hashedPassword, username)
	if err != nil {
		utils.Log.Error("Failed to change password: ", err)
		return err
	}

	utils.Log.Info("Password changed successfully for user: ", username)
	return nil
}

// SetPermission updates the permission level for an existing user
func SetPermission(username, permission string) error {
	if !isValidPermission(permission) {
		err := errors.New("invalid permission")
		utils.Log.Error("SetPermission failed: ", err)
		return err
	}

	_, err := db.Exec(`UPDATE users SET permission = ? WHERE username = ?`, permission, username)
	if err != nil {
		utils.Log.Error("Failed to set permission: ", err)
		return err
	}

	utils.Log.Info("Permission updated successfully for user: ", username)
	return nil
}

// ListUsers returns all users and their permissions
func ListUsers() ([]map[string]string, error) {
	rows, err := db.Query(`SELECT username, permission FROM users`)
	if err != nil {
		utils.Log.Error("Failed to list users: ", err)
		return nil, err
	}
	defer rows.Close()

	var users []map[string]string
	for rows.Next() {
		var username, permission string
		if err := rows.Scan(&username, &permission); err != nil {
			utils.Log.Error("Failed to scan user: ", err)
			return nil, err
		}
		users = append(users, map[string]string{"username": username, "permission": permission})
	}

	utils.Log.Info("Listed users successfully")
	return users, nil
}

// GetPermission returns the permission of a given user
func GetPermission(username string) (string, error) {
	var permission string
	err := db.QueryRow(`SELECT permission FROM users WHERE username = ?`, username).Scan(&permission)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.Log.Warn("User not found: ", username)
			return "", errors.New("user not found")
		}
		utils.Log.Error("Failed to get permissions: ", err)
		return "", err
	}

	utils.Log.Info("Fetched permissions successfully for user: ", username)
	return permission, nil
}

// HasPermission checks if a user has a specific permission
func HasPermission(username, permission string) (bool, error) {
	currentPermission, err := GetPermission(username)
	if err != nil {
		return false, err
	}

	return currentPermission == permission, nil
}

// isValidPermission checks if the provided permission is valid
func isValidPermission(permission string) bool {
	return permission == SystemAdmin || permission == Admin || permission == Insight
}

func CreateDefaultUserIfNecessary() {
	// Create default user named "admin" with password "admin" with system-admin permissions if no other users exist
	users, err := ListUsers()
	if err != nil {
		utils.Log.Error(err)
		utils.Log.Panic("Could not access users, so wasn't able to check for default user!")
		panic(err)
	}

	if len(users) == 0 {
		// No users created (yet), create default user:
		err := CreateUser("admin", "admin", SystemAdmin)
		if err == nil {
			utils.Log.Info("Successfully created new default user 'admin' with the password 'admin'!")
		} else {
			utils.Log.Error("Could not create default user with name 'admin', will try to create a user based on uuid!")
			defaultUsername := "admin-" + uuid.NewString()
			err := CreateUser(defaultUsername, "admin", SystemAdmin)
			if err == nil {
				utils.Log.Info("Successfully created new default user '" + defaultUsername + "' with the password 'admin'!")
			}
		}
	} else {
		// Make sure at least one user has system-admin permissions, else create default user

		foundSysAdmin := false

		for i := 0; i < len(users); i++ {
			currentUser := users[i]
			if currentUser["permission"] == SystemAdmin {
				foundSysAdmin = true
			}
		}

		if foundSysAdmin {
			utils.Log.Info("No need to create new default user, as at least one user with system-admin permissions already exists!")
		} else {
			err := CreateUser("admin", "admin", SystemAdmin)
			if err == nil {
				utils.Log.Info("Successfully created new default user 'admin' with the password 'admin' because no user with system-admin permissions was found!")
			} else {
				utils.Log.Error("Could not create default user with name 'admin', will try to create a user based on uuid!")
				defaultUsername := "admin-" + uuid.NewString()
				err := CreateUser(defaultUsername, "admin", SystemAdmin)
				if err == nil {
					utils.Log.Info("Successfully created new default user '" + defaultUsername + "' with the password 'admin' because no user with system-admin permissions was found!")
				}
			}
		}
	}
}
