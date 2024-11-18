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

	// Start Server
	go web.StartServer()

	// Create default user named "admin" with password "admin" with system-admin permissions if no other users exist
	users, err := auth.ListUsers()
	if err != nil {
		utils.Log.Error(err)
		utils.Log.Panic("Could not access users, so wasn't able to check for default user!")
		panic(err)
	}

	if len(users) == 0 {
		// No users created (yet), create default user:
		auth.CreateUser("admin", "admin", auth.SystemAdmin)
		utils.Log.Info("Successfully created new default user 'admin' with the password 'admin'!")
	} else {
		utils.Log.Info("No need to create new default user, as at least one user already exists!")
	}
}

// TODO: Protect API by user auth
// TODO: Implement Login in frontend
// TODO: User Creation UI
