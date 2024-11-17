package routes

import (
	"Dampfer/utils"
	"encoding/json"
	"net/http"
	"os/exec"
	"strings"
)

func ListAllContainers(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("docker", "ps", "-a", "--no-trunc", "--format", "{{json .}}")
	output, err := cmd.Output()
	if err != nil {
		utils.Log.Warn("Failed to list all containers: ", err)
		http.Error(w, "Failed to list all containers", http.StatusInternalServerError)
		return
	}

	var containers []map[string]interface{}
	for _, line := range strings.Split(strings.TrimSpace(string(output)), "\n") {
		var container map[string]interface{}
		if err := json.Unmarshal([]byte(line), &container); err == nil {
			containers = append(containers, container)
		} else {
			utils.Log.Warn("Failed to parse container JSON line: ", line, "Error: ", err)
		}
	}

	json.NewEncoder(w).Encode(containers)
}
