package utils

import (
	"encoding/json"
	"strings"
)

func ParseDockerCmdToJson(cmdOutput []byte) []map[string]interface{} {
	// Basic assumption: Docker outputs every item to list (eg. a single container) as a separate line and a json object:
	// {container1...}
	// {container2...}
	// If no container was found the output will just be empty new lines

	// Stare all object like containers in here
	var listed []map[string]interface{} = []map[string]interface{}{}

	// Split by new lines and
	for _, line := range strings.Split(strings.TrimSpace(string(cmdOutput)), "\n") {
		var listItem map[string]interface{}

		// Catch empty lines
		if line == "" {
			continue
		}

		if err := json.Unmarshal([]byte(line), &listItem); err == nil {
			listed = append(listed, listItem)
		} else {
			Log.Warn("Failed to parse container JSON line: ", line, "Error: ", err)
		}
	}

	return listed
}
