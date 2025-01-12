package api

import (
	"Dampfer/utils"
	"database/sql"
)

func GetPasswordChangeCount(username string) (int, error) {
	var changeCount int

	err := db.QueryRow(`SELECT change_count FROM password_changes WHERE username = ?`, username).Scan(&changeCount)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.Log.Warn("No password change count found for user: ", username)
			return 0, nil // Return 0 if the user does not exist in the table
		}
		utils.Log.Error("Failed to retrieve password change count: ", err)
		return 0, err
	}

	utils.Log.Info("Password change count retrieved successfully for user: ", username)
	return changeCount, nil
}
