package docker

import (
	"Dampfer/utils"
	"os/exec"
)

func ListAllContainers() ([]map[string]interface{}, error) {
	cmd := exec.Command("docker", "ps", "-a", "--no-trunc", "--format", "{{json .}}")
	output, err := cmd.Output()
	if err != nil {
		utils.Log.Warn("Failed to list all containers: ", err)
		return nil, err
	}

	return utils.ParseDockerCmdToJson(output), nil
}

func ListRunningContainers() ([]map[string]interface{}, error) {
	cmd := exec.Command("docker", "ps", "--no-trunc", "--format", "{{json .}}")
	output, err := cmd.Output()
	if err != nil {
		utils.Log.Warn("Failed to list running containers: ", err)
		return nil, err
	}

	return utils.ParseDockerCmdToJson(output), nil
}

func ListImages() ([]map[string]interface{}, error) {
	cmd := exec.Command("docker", "images", "--format", "{{json .}}")
	output, err := cmd.Output()
	if err != nil {
		utils.Log.Warn("Failed to list images: ", err)
		return nil, err
	}

	return utils.ParseDockerCmdToJson(output), nil
}

func ListVolumes() ([]map[string]interface{}, error) {
	cmd := exec.Command("docker", "volume", "ls", "--format", "{{json .}}")
	output, err := cmd.Output()
	if err != nil {
		utils.Log.Warn("Failed to list volumes: ", err)
		return nil, err
	}

	return utils.ParseDockerCmdToJson(output), nil
}
