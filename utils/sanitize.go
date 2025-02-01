package utils

import "strings"

// SanitizeCommandPortion wraps the string in quotes and escapes quotes and backslashes
func SanitizeCommandPortion(portion string) string {
	// Escape backslashes first, then escape double quotes
	escaped := strings.ReplaceAll(portion, `\`, `\\`)
	escaped = strings.ReplaceAll(escaped, `"`, `\"`)

	// Return the portion wrapped in quotes
	return `"` + escaped + `"`
}
