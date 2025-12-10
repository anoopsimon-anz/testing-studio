package handlers

import (
	"encoding/json"
	"net/http"
	"os/exec"
)

type DockerStatusResponse struct {
	Running bool `json:"running"`
}

func HandleDockerStatus(w http.ResponseWriter, r *http.Request) {
	// Check if Docker is running by executing 'docker info'
	cmd := exec.Command("docker", "info")
	err := cmd.Run()

	response := DockerStatusResponse{
		Running: err == nil,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
