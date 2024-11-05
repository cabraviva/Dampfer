package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"os/exec"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

// Embed files in the svelte/public directory
//
//go:embed svelte/dist/*
var embeddedFiles embed.FS

// stores all registered endpoints
var registeredEndpoints []string

// Register Logger
var log = logrus.New()

func main() {
	// Logging
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	log.SetLevel(logrus.InfoLevel)

	// Check if Docker is installed and log result
	if installed, installMsg := dockerInstalled(); installed {
		log.Info("Docker installation status: Installed")
	} else {
		log.Warn("Docker installation status: Not installed - " + installMsg)
	}

	// Check if Docker daemon is running and log result
	if running, daemonMsg := dockerDaemonRunning(); running {
		log.Info("Docker daemon status: Running")
	} else {
		log.Error("Docker daemon status: Not running - " + daemonMsg)
	}

	// Serve embedded files as an HTTP file system
	svelteFiles, _ := fs.Sub(embeddedFiles, "svelte/dist")
	fs := http.FileServer(http.FS(svelteFiles))
	http.Handle("/", fs)

	// Register API endpoints with their handler functions and allowed HTTP methods
	registerAPI("/api/docker-ready", API_dockerReady, http.MethodGet)
	registerAPI("/api/endpoints", API_listEndpoints, http.MethodGet)

	log.Info("Server started on port 13777")
	log.Fatal(http.ListenAndServe(":13777", nil))

}

// registerAPI registers an API endpoint with a specific function and allowed method,
// and adds it to the list of endpoints for listing at /api/endpoints.
func registerAPI(path string, handler func(http.ResponseWriter, *http.Request), method string) {
	// Record the endpoint in the format "METHOD /path"
	registeredEndpoints = append(registeredEndpoints, fmt.Sprintf("%s %s", method, path))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		handler(w, r)
	})
}

// API_listEndpoints lists all registered API endpoints in plain text
func API_listEndpoints(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(strings.Join(registeredEndpoints, "\n")))
}

// API_dockerReady is the handler function for the /api/docker-ready endpoint
func API_dockerReady(w http.ResponseWriter, r *http.Request) {
	response := dockerReady()
	json.NewEncoder(w).Encode(response)
}

// dockerReady checks if Docker is installed and if the daemon is running
func dockerReady() map[string]interface{} {
	installed, installMsg := dockerInstalled()
	if !installed {
		return map[string]interface{}{
			"ready":         false,
			"installed":     false,
			"daemonRunning": false,
			"msg":           installMsg,
		}
	}

	daemonRunning, daemonMsg := dockerDaemonRunning()
	if !daemonRunning {
		return map[string]interface{}{
			"ready":         false,
			"installed":     true,
			"daemonRunning": false,
			"msg":           daemonMsg,
		}
	}

	return map[string]interface{}{
		"ready":         true,
		"installed":     true,
		"daemonRunning": true,
		"msg":           "Docker is ready and running",
	}
}

// dockerInstalled checks if Docker is installed on the system by running a version check
func dockerInstalled() (bool, string) {
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
func dockerDaemonRunning() (bool, string) {
	cmd := exec.Command("docker", "info")
	err := cmd.Run()
	if err != nil {
		return false, "Docker Daemon is not running"
	}
	return true, ""
}
