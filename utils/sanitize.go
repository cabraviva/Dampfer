package utils

import (
	"regexp"
)

// Allowed: The regular expression ^[a-zA-Z0-9_\-\.]+$ ensures that the input contains only alphanumeric characters, underscores, hyphens, and dots.
func SanitizeCommandPortion(portion string) string {
	// Define a regular expression for allowed characters (alphanumeric, _, -, .)
	re := regexp.MustCompile(`^[a-zA-Z0-9_\-\.]+$`)

	// Check if the input matches the allowed pattern
	if !re.MatchString(portion) {
		Log.Warn("Some bad command portion was detected:" + portion + "\nIt will be ignored")
		return ""
	}

	// Return the validated input
	return portion
}
