package routes

import (
	"Dampfer/api"
	"encoding/json"
	"fmt"
	"net/http"
)

func MyPwChanges(w http.ResponseWriter, r *http.Request, username string) {
	// return the number of times the password was changed
	changeCount, err := api.GetPasswordChangeCount(username)
	if err != nil {
		fmt.Println("Error retrieving password change count:", err)
		json.NewEncoder(w).Encode(0)
	} else {
		fmt.Printf("Password change count for %s: %d\n", username, changeCount)
		json.NewEncoder(w).Encode(changeCount)
	}
}
