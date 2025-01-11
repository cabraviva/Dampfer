package web

import (
	"Dampfer/api"
	"Dampfer/auth"
	"Dampfer/routes"
	"net/http"
)

func RegisterEndpoints() {
	api.Register("/api/endpoints", api.ListEndpoints, http.MethodGet, true, auth.Insight)

	api.Register("/login", routes.Login, http.MethodPost, false, "")
	api.Register("/api/whoami", routes.Whoami, http.MethodGet, true, auth.Insight)

	api.Register("/api/users/ls", routes.ListUsers, http.MethodGet, true, auth.SystemAdmin)
	api.Register("/api/users/delete", routes.DeleteUser, http.MethodPost, true, auth.SystemAdmin)
	api.Register("/api/users/create", routes.CreateUser, http.MethodPost, true, auth.SystemAdmin)
	api.Register("/api/users/set-permission", routes.SetUserPermission, http.MethodPost, true, auth.SystemAdmin)
	api.Register("/api/users/set-password-sysadmin", routes.SetUserPassword, http.MethodPost, true, auth.SystemAdmin)

	api.Register("/api/docker/ready", routes.DockerReady, http.MethodGet, true, auth.Insight)
	api.Register("/api/docker/container/running/list", routes.ListRunningContainers, http.MethodGet, true, auth.Insight)
	api.Register("/api/docker/container/all/list", routes.ListAllContainers, http.MethodGet, true, auth.Insight)
	api.Register("/api/docker/image/list", routes.ListImages, http.MethodGet, true, auth.Insight)
	api.Register("/api/docker/volume/list", routes.ListVolumes, http.MethodGet, true, auth.Insight)
}
