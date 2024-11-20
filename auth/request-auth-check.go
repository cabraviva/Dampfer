package auth

import (
	"Dampfer/utils"
	"net/http"
	"strings"
)

func isPermissionSufficient(permission string, minimumPermission string) bool {

	if minimumPermission == "" {
		// No permissions required
		return true
	}

	if permission == "" {
		// User has no permissions
		return false
	}

	if !isValidPermission(permission) || !isValidPermission(minimumPermission) {
		utils.Log.Warn("Permission '" + permission + "' or permission '" + minimumPermission + "' is not a valid permission. Request not authorized.")
		return false
	}

	if minimumPermission == Insight {
		// Always fulfilled as any other invalid permission is guarded off above
		return true
	}

	if minimumPermission == Admin {
		return permission == Admin || permission == SystemAdmin
	}

	if minimumPermission == SystemAdmin {
		return permission == SystemAdmin
	}

	utils.Log.Warn("Code reached that should be unreachable at auth.isPermissionSufficient() in file Dampfer/auth/request-auth-check.go! Permission: ", permission, " Minimum permission: ", minimumPermission)
	return false
}

// Returns bool isAuth, string erMsg, string username
func IsRequestAuthorized(w http.ResponseWriter, r *http.Request, minimumPermissions string) (bool, string, string) {
	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return false, "Authorization header is missing", ""
	}

	// Check if the Authorization header starts with "Bearer "
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return false, "Invalid Authorization scheme", ""
	}

	// Extract the token part
	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		return false, "Bearer token is missing", ""
	}

	// Check if token is valid
	isValid, username, permission, err := ValidateToken(token)

	if !isValid || err != nil {
		// Invalid token
		return false, err.Error(), ""
	}

	if isPermissionSufficient(permission, minimumPermissions) {
		return true, "", username
	} else {
		return false, "Permissions insufficient", username
	}
}
