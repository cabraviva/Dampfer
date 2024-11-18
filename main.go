package main

import (
	"net/http"

	"Dampfer/docker"
	"Dampfer/utils"
	"Dampfer/web"
)

var svelteFS http.Handler // Shared across dev.go and prod.go

func main() {
	// Logging
	utils.InitLogger()

	// Check if Docker is installed and log result
	docker.DoStartupInstallChecks()

	// Init HTTP Server
	web.InitServer(svelteFS)

	// Register API endpoints
	web.RegisterEndpoints()

	// Start Server
	web.StartServer()
}

// TODO: Protect API by user auth
// TODO: Implement Login in frontend
// TODO: User Creation UI
