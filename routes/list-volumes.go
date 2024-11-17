package routes

import (
	"Dampfer/utils"
	"encoding/json"
	"net/http"
	"os/exec"
	"strings"
)

func ListVolumes(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("docker", "volume", "ls", "--format", "{{json .}}")
	output, err := cmd.Output()
	if err != nil {
		utils.Log.Warn("Failed to list volumes: ", err)
		http.Error(w, "Failed to list volumes", http.StatusInternalServerError)
		return
	}

	var volumes []map[string]interface{}
	for _, line := range strings.Split(strings.TrimSpace(string(output)), "\n") {
		var volume map[string]interface{}
		if err := json.Unmarshal([]byte(line), &volume); err == nil {
			volumes = append(volumes, volume)
		} else {
			utils.Log.Warn("Failed to parse volume JSON line: ", line, "Error: ", err)
		}
	}

	json.NewEncoder(w).Encode(volumes)
}
