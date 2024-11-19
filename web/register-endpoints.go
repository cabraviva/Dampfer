package web

import (
	"Dampfer/api"
	"Dampfer/routes"
	"net/http"
)

func RegisterEndpoints() {
	api.Register("/api/endpoints", api.ListEndpoints, http.MethodGet)

	api.Register("/api/docker-ready", routes.DockerReady, http.MethodGet)
	api.Register("/api/docker-running-containers", routes.ListRunningContainers, http.MethodGet)
	api.Register("/api/docker-all-containers", routes.ListAllContainers, http.MethodGet)
	api.Register("/api/docker-images", routes.ListImages, http.MethodGet)
	api.Register("/api/docker-volumes", routes.ListVolumes, http.MethodGet)
}
