package utils

import (
	"os/exec"
	"runtime"
)

func DoStartupInstallChecks() {
	if installed, installMsg := DockerInstalled(); installed {
		Log.Info("Docker installation status: Installed")
	} else {
		Log.Error("Docker installation status: Not installed - " + installMsg)
	}

	// Check if Docker daemon is running and log result
	if running, daemonMsg := DockerDaemonRunning(); running {
		Log.Info("Docker daemon status: Running")
	} else {
		Log.Error("Docker daemon status: Not running - " + daemonMsg)
	}

	// Check if Docker Compose is installed and log result
	if composeInstalled, composeMsg := DockerComposeInstalled(); composeInstalled {
		Log.Info("Docker Compose installation status: Installed")
	} else {
		Log.Error("Docker Compose installation status: Not installed - " + composeMsg)
	}
}

// dockerComposeInstalled checks if Docker Compose is installed on the system
func DockerComposeInstalled() (bool, string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("docker-compose", "--version")
	case "darwin", "linux":
		cmd = exec.Command("which", "docker-compose")
	}

	err := cmd.Run()
	if err != nil {
		return false, "Docker Compose does not seem to be installed or added to PATH on your system"
	}
	return true, ""
}

// dockerInstalled checks if Docker is installed on the system by running a version check
func DockerInstalled() (bool, string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("docker", "--version")
	case "darwin", "linux":
		cmd = exec.Command("which", "docker")
	}

	err := cmd.Run()
	if err != nil {
		return false, "Docker does not seem to be installed or added to PATH on your system"
	}
	return true, ""
}

// dockerDaemonRunning checks if the Docker daemon is running by executing a simple docker command
func DockerDaemonRunning() (bool, string) {
	cmd := exec.Command("docker", "info")
	err := cmd.Run()
	if err != nil {
		return false, "Docker Daemon is not running"
	}
	return true, ""
}
