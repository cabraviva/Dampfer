package docker

import (
	"Dampfer/database"
	"Dampfer/icongen"
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
	// Execute Docker command to list images in JSON format
	cmd := exec.Command("docker", "images", "--format", "{{json .}}")
	output, err := cmd.Output()
	if err != nil {
		utils.Log.Warn("Failed to list images: ", err)
		return nil, err
	}

	// Use utils.ParseDockerCmdToJson to parse the Docker images output
	images := utils.ParseDockerCmdToJson(output)

	// Iterate over each image
	for _, image := range images {
		imageID := image["ID"].(string)
		repository := image["Repository"].(string)

		// Check if the image is already stored in the database
		var count int
		err := database.DB.QueryRow("SELECT COUNT(*) FROM icon_cache WHERE id = ?", imageID).Scan(&count)
		if err != nil {
			utils.Log.Warn("Error checking icon in database: ", err)
			continue // Log error but continue processing other images
		}

		// If the image is not found in the database, attempt to download and save it
		if count == 0 {
			// Assuming GetImagesFor returns URLs for images based on repository name
			searchResults, err := icongen.GetImagesFor(repository)
			if err != nil || len(searchResults) == 0 {
				utils.Log.Warn("No icons found for repository: ", repository)
				continue
			}

			// Attempt to download and save the image to the database
			err = icongen.DownloadImageToDB(imageID, searchResults[0]) // Assuming the first result is the desired image
			if err != nil {
				utils.Log.Warn("Failed to download and save icon for docker image: ", repository, err)
				continue
			}

			utils.Log.Info("Successfully downloaded and saved icon for docker image: ", repository, imageID)
		}
	}

	return images, nil
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
