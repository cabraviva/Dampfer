package docker

import (
	"os/exec"
)

type ComposeVersion string

const (
	ComposeVersionV1       ComposeVersion = "v1"
	ComposeVersionV2       ComposeVersion = "v2"
	ComposeVersionNotFound ComposeVersion = "NOT_FOUND"
)

// Checks whether docker-compose (v1) is installed
func IsComposeV1Installed() bool {
	cmd := exec.Command("docker-compose", "--version")

	err := cmd.Run()

	return err == nil
}

// Checks whether docker compose (v2) is installed
func IsComposeV2Installed() bool {
	cmd := exec.Command("docker", "compose", "version")

	err := cmd.Run()

	return err == nil
}

// Returns installed compose version ("v1" or "v2") or if not installed "NOT_FOUND"
func GetComposeVersion() ComposeVersion {
	// Prefer v2
	if IsComposeV2Installed() {
		return ComposeVersionV1
	} else if IsComposeV1Installed() {
		return ComposeVersionV2
	} else {
		return ComposeVersionNotFound
	}
}

// Returns true if some compose version is installed
func IsComposeInstalled() bool {
	return GetComposeVersion() != "NOT_FOUND"
}
