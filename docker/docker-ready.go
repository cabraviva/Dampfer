package docker

import "Dampfer/utils"

// dockerReady checks if Docker is installed and if the daemon is running
func DockerReady() map[string]interface{} {
	installed, installMsg := utils.DockerInstalled()
	if !installed {
		return map[string]interface{}{
			"ready":            false,
			"installed":        false,
			"daemonRunning":    false,
			"composeInstalled": false,
			"msg":              installMsg,
		}
	}

	daemonRunning, daemonMsg := utils.DockerDaemonRunning()
	if !daemonRunning {
		return map[string]interface{}{
			"ready":            false,
			"installed":        true,
			"daemonRunning":    false,
			"composeInstalled": false,
			"msg":              daemonMsg,
		}
	}

	composeInstalled, composeMsg := utils.DockerComposeInstalled()
	if !composeInstalled {
		return map[string]interface{}{
			"ready":            false,
			"installed":        true,
			"daemonRunning":    true,
			"composeInstalled": false,
			"msg":              composeMsg,
		}
	}

	return map[string]interface{}{
		"ready":            true,
		"installed":        true,
		"daemonRunning":    true,
		"composeInstalled": true,
		"msg":              "Docker and Docker Compose are ready and running",
	}
}
