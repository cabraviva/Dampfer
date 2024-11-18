package docker

import (
	"Dampfer/utils"
	"os/exec"
)

// Performs some install checks on startup and logs results
func DoStartupInstallChecks() {
	readyStatus := IsReady()

	if readyStatus.Installed {
		utils.Log.Info("Docker installation status: Installed")
	} else {
		utils.Log.Error("Docker installation status: Not installed - " + readyStatus.Msg)
	}

	// Check if Docker daemon is running and log result
	if readyStatus.DaemonRunning {
		utils.Log.Info("Docker daemon status: Running")
	} else {
		utils.Log.Error("Docker daemon status: Not running - " + readyStatus.Msg)
	}

	// Check if Docker Compose is installed and log result
	if readyStatus.DaemonRunning {
		utils.Log.Info("Docker Compose installation status: Installed; Version: " + readyStatus.ComposeVersion)
	} else {
		utils.Log.Error("Docker Compose installation status: Not installed - " + readyStatus.Msg)
	}
}

// dockerInstalled checks if Docker is installed on the system by running a version check
func IsDockerInstalled() bool {
	cmd := exec.Command("docker", "--version")
	err := cmd.Run()
	return err == nil
}

// dockerDaemonRunning checks if the Docker daemon is running by executing a simple docker command
func IsDockerDaemonRunning() bool {
	cmd := exec.Command("docker", "info")
	err := cmd.Run()
	return err == nil
}
