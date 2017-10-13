package entities

import (
	"sync"
)

const (
	DockerCommand      = "docker"
	DockerBuild        = "build"
	DockerBuildForceRM = "--force-rm" // Force deletion of the temporary image.
	DockerBuildTag     = "-t"         // Give a tag to the built image.
	DockerRun          = "run"
	DockerRunRM        = "--rm"   // Delete container upon completion of run.
	DockerRunName      = "--name" // Name the container.
	DockerStop         = "stop"   // Stop the named container. Used if the container is killed prematurely.
	DockerRM           = "rm"     // Destroy the named container. Used if the container is killed prematurely.
	DockerScriptName   = "script.sh"
	DockerFile         = "Dockerfile"
	ContainerPrefix    = "executor_" // Namespace containers created by the executor.
)

const (
	NotFound             = iota
	Queued               = iota
	MakingBuildDirectory = iota
	CopyingDockerfile    = iota
	CopyingScript        = iota
	RunningTest          = iota
	BuildingImage        = iota
	Starting             = iota
	StartingDocker       = iota
	DockerStarted        = iota
	SendingScripts       = iota
	ScriptsSent          = iota
	ExecutingScripts     = iota
	ResultsReceived      = iota
	DestroyingDocker     = iota
	DockerDestroyed      = iota
	Done                 = iota
	Failed               = iota
)

var stateMap = map[int]string{
	NotFound:             "Not found.",
	Queued:               "awaiting execution.",
	MakingBuildDirectory: "Making build directory",
	BuildingImage:        "Building image.",
	Starting:             "starting.",
	StartingDocker:       "Docker is initializing.",
	DockerStarted:        "Docker has started.",
	SendingScripts:       "Sending artifacts to the Docker container.",
	ScriptsSent:          "Artifacts have been sent to the Docker container.",
	ExecutingScripts:     "is executing.",
	DestroyingDocker:     "Docker is being torn down.",
	DockerDestroyed:      "Docker has been torn down.",
	Done:                 "Done.",
}

type ContainerRunEntity struct {
	ID                   int    `json:"id"`
	EnvironmentVariables string `json:"environmentVariables"`
	Dockerfile           string `json:"dockerfile"`
	Script               string `json:"script"`
	Cancel               chan uint8
	state                int
	mutex                sync.RWMutex
}
