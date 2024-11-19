package main

import (
	"Dampfer/auth"
	"Dampfer/docker" // Import the restartCrash package
	"Dampfer/utils"
	"Dampfer/web"

	_ "Dampfer/database" // Necessary to init database
	"net/http"
)

var svelteFS http.Handler // Shared across dev.go and prod.go

var isDEV bool // Set by dev.go or prod.go

func main() {
	// Logging
	utils.InitLogger()

	if isDEV {
		utils.Log.Warn("Using DEVELOPMENT mode. All files will be served from the local FS and aren't embedded.")
	} else {
		utils.Log.Info("Starting in PRODUCTION mode. Webserver will be using embedded files. That's a good sign 👍")
	}

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
