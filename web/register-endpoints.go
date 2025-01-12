package web

import (
	"Dampfer/api"
	"Dampfer/auth"
	"Dampfer/routes"
	"net/http"
)

func RegisterEndpoints() {
	// List all endpoints
	api.Register("/api/endpoints", api.ListEndpoints, http.MethodGet, true, auth.SystemAdmin)

	//  Login
	api.Register("/login", routes.Login, http.MethodPost, false, "")

	// User Management
	api.Register("/api/users/ls", routes.ListUsers, http.MethodGet, true, auth.SystemAdmin)
	api.Register("/api/users/delete", routes.DeleteUser, http.MethodPost, true, auth.SystemAdmin)
	api.Register("/api/users/create", routes.CreateUser, http.MethodPost, true, auth.SystemAdmin)
	api.Register("/api/users/set-permission", routes.SetUserPermission, http.MethodPost, true, auth.SystemAdmin)
	api.Register("/api/users/set-password-sysadmin", routes.SetUserPassword, http.MethodPost, true, auth.SystemAdmin)

	// Me (current user)
	api.Register("/api/whoami", routes.Whoami, http.MethodGet, true, auth.Insight)
	api.Register("/api/me/change-password", routes.ChangeMyPassword, http.MethodPost, true, auth.Insight)
	api.Register("/api/me/how-often-was-pw-changed", routes.MyPwChanges, http.MethodGet, true, auth.Insight)

	// ################# Docker #################
	// Docker Ready
	api.Register("/api/docker/ready", routes.DockerReady, http.MethodGet, true, auth.Insight)

	// Containers
	api.Register("/api/docker/container/running/list", routes.ListRunningContainers, http.MethodGet, true, auth.Insight)
	api.Register("/api/docker/container/all/list", routes.ListAllContainers, http.MethodGet, true, auth.Insight)

	// Images
	api.Register("/api/docker/image/list", routes.ListImages, http.MethodGet, true, auth.Insight)

	// Volumes
	api.Register("/api/docker/volume/list", routes.ListVolumes, http.MethodGet, true, auth.Insight)
}
