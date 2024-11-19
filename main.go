package main

import (
	"net/http"

	"Dampfer/auth"
	"Dampfer/docker"
	"Dampfer/utils"
	"Dampfer/web"

	_ "Dampfer/database" // Import to init db
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

	// Create default user named "admin" with password "admin" with system-admin permissions if no other users exist
	auth.CreateDefaultUserIfNecessary()

	// Start Server
	go web.StartServer()
}

// TODO: Protect API by user auth
// TODO: Implement Login in frontend
// TODO: User Creation UI
