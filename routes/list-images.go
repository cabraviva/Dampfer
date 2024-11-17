package routes

import (
	"Dampfer/utils"
	"encoding/json"
	"net/http"
	"os/exec"
	"strings"
)

func ListImages(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("docker", "images", "--format", "{{json .}}")
	output, err := cmd.Output()
	if err != nil {
		utils.Log.Warn("Failed to list images: ", err)
		http.Error(w, "Failed to list images", http.StatusInternalServerError)
		return
	}

	var images []map[string]interface{}
	for _, line := range strings.Split(strings.TrimSpace(string(output)), "\n") {
		var image map[string]interface{}
		if err := json.Unmarshal([]byte(line), &image); err == nil {
			images = append(images, image)
		} else {
			utils.Log.Warn("Failed to parse image JSON line: ", line, "Error: ", err)
		}
	}

	json.NewEncoder(w).Encode(images)
}
