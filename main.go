package main

import (
	"net/http"

	"Dampfer/utils"
	"Dampfer/web"
)

var svelteFS http.Handler // Shared across dev.go and prod.go

func main() {
	// Logging
	utils.InitLogger()

	// Check if Docker is installed and log result
	utils.DoStartupInstallChecks()

	// Init HTTP Server
	web.InitServer(svelteFS)

	// Register API endpoints
	web.RegisterEndpoints()

	// Start Server
	web.StartServer()
}

// TODO: Most endpoints don't have corresponding docker.Function handler atm
