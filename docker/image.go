package docker

import (
	"Dampfer/utils"
	"os/exec"
	"strings"
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

func ImageRm(repository string, tag string) (bool, error) {
	cmd := exec.Command("docker", "image", "rm", repository+":"+tag)
	output, err := cmd.Output()
	if err != nil {
		utils.Log.Warn("Failed to delete image '"+repository+":"+tag+"': ", err)
		return false, err
	}

	var success bool
	if !cmd.ProcessState.Success() {
		utils.Log.Warn("Failed to delete image '"+repository+":"+tag+"': ", "Cmd exited with non 0 code")
		return false, nil
	}

	if strings.Contains(string(output), "Untagged") {
		success = true
	} else {
		success = false
		utils.Log.Warn("Failed to delete image '"+repository+":"+tag+"': ", "Output does not contain 'Untagged': ", string(output))
	}

	return success, nil
}
