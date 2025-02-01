package docker

import (
	"Dampfer/utils"
	"os/exec"
)

func ImageInspect(id string) ([]map[string]interface{}, error) {
	cmd := exec.Command("docker", "image", "inspect", utils.SanitizeCommandPortion(id), "--format", "{{json .}}")
	output, err := cmd.Output()
	if err != nil {
		utils.Log.Warn("Failed to inspect image '"+id+"': ", err)
		return nil, err
	}

	return utils.ParseDockerCmdToJson(output), nil
}
