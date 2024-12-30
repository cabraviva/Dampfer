package docker

type DockerReadyStatus struct {
	Ready            bool
	Installed        bool
	DaemonRunning    bool
	ComposeInstalled bool
	ComposeVersion   ComposeVersion
	Msg              string
}

func IsReady() DockerReadyStatus {
	dockerInstalled, daemonRunning, composeInstalled, composeVersion := IsDockerInstalled(), IsDockerDaemonRunning(), IsComposeInstalled(), GetComposeVersion()

	var msg string

	if dockerInstalled && daemonRunning && composeInstalled {
		msg = "Docker and Compose are installed and running!"
	} else if dockerInstalled && composeInstalled && !daemonRunning {
		msg = "Docker and Compose are installed but Docker Daemon is not running. Make sure docker is configured to start with the system!"
	} else if dockerInstalled && !composeInstalled && daemonRunning {
		msg = "Docker is installed and running, but Dampfer wasn't able to find Compose on your system. Make sure you have either" +
			" \"docker-compose\" or \"docker compose\" installed and added to PATH!"
	} else if dockerInstalled && !composeInstalled && !daemonRunning {
		msg = "Docker is installed but the docker daemon is not running and Dampfer wasn't able to find Compose on your system. " +
			"Make sure you have either \"docker-compose\" or \"docker compose\" installed and added to PATH! Also make sure docker is configured to start with the system!"
	} else if !dockerInstalled && composeInstalled {
		msg = "Docker does not seem to be installed and/or added to PATH on your system!"
	} else if !dockerInstalled && !composeInstalled {
		msg = "Seems like Docker and Docker Compose aren't installed and/or added to PATH on your system! Most of the time both get's fixed by just (re-)installing Docker!"
	} else {
		msg = "Dampfer detected some kind of problem with your Docker and Docker Compose setup. Either the Docker Daemon isn't running, Docker is not installed or Docker Compose was not found on your system!"
	}

	return DockerReadyStatus{
		Ready:            dockerInstalled && daemonRunning && composeInstalled,
		Installed:        dockerInstalled,
		DaemonRunning:    daemonRunning,
		ComposeInstalled: composeInstalled,
		ComposeVersion:   composeVersion,
		Msg:              msg,
	}
}
