package controllers

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"udit/restart-container/models"
)

func RestartContainer(dockerContainerModel models.DockerContainerModel) error {
	// Constants from your bash script
	CONTAINER_NAME := dockerContainerModel.ContainerName
	IMAGE_NAME := dockerContainerModel.ImageName
	PORT := dockerContainerModel.Port

	// Log "running"
	log.Println("running")

	// Stop the container (ignore errors)
	stopCmd := exec.Command("docker", "stop", CONTAINER_NAME)
	stopCmd.Run()

	// Remove the container (ignore errors)
	rmCmd := exec.Command("docker", "rm", CONTAINER_NAME)
	rmCmd.Run()

	// Run new container with latest image
	fmt.Println("Restarting container with latest image...")
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	envPath := filepath.Join(homeDir, "secrets", ".env")

	if dockerContainerModel.ShouldPull {

		runCmd := exec.Command("docker", "run", "-d", "--pull=always", "--quiet",
			"--name", CONTAINER_NAME,
			"-v", envPath+":/app/.env",
			"--env", "PORT="+PORT,
			"--publish", PORT+":"+PORT,
			IMAGE_NAME)

		output, err := runCmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("error restarting container: %v, output: %s", err, output)
		}

	} else {

		runCmd := exec.Command("docker", "run", "-d",
			"--name", CONTAINER_NAME,
			"--env", "PORT="+PORT,
			"--publish", PORT+":"+PORT,
			IMAGE_NAME)

		output, err := runCmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("error restarting container: %v, output: %s", err, output)
		}

	}

	return nil
}
