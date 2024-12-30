package routes

import (
	"Dampfer/auth"
	"Dampfer/utils"
	"encoding/json"
	"net/http"
)

type WhoamiType struct {
	Username    string `json:"username"`
	Permission  string `json:"permission"`
	Insight     bool   `json:"insight"`
	Admin       bool   `json:"admin"`
	SystemAdmin bool   `json:"systemAdmin"`
}

func Whoami(w http.ResponseWriter, r *http.Request, username string) {
	highestPermission, error := auth.GetPermission(username)

	insight := false
	admin := false
	systemAdmin := false

	if error != nil {
		returnValue := WhoamiType{
			Username:    username,
			Permission:  "unknown",
			Insight:     false,
			Admin:       false,
			SystemAdmin: false,
		}

		json.NewEncoder(w).Encode(returnValue)
		return
	}

	if highestPermission == auth.Insight {
		insight = true
	} else if highestPermission == auth.Admin {
		insight = true
		admin = true
	} else if highestPermission == auth.SystemAdmin {
		insight = true
		admin = true
		systemAdmin = true
	} else {
		utils.Log.Error("User " + username + " has invalid permissions. Highest permission is: " + highestPermission)
	}

	returnValue := WhoamiType{
		Username:    username,
		Permission:  highestPermission,
		Insight:     insight,
		Admin:       admin,
		SystemAdmin: systemAdmin,
	}

	json.NewEncoder(w).Encode(returnValue)
}
