package handlers

import (
	"fmt"
	"log"
	"net/http"
	"udit/restart-container/controllers"
	"udit/restart-container/models"

	"github.com/gin-gonic/gin"
)

func RestartDockerHandler(router *gin.Engine) {

	var routePrefix string = "/API/RestartDockerContainer"

	router.POST(routePrefix+"/apiPadhai", apiPadhaiWebhookHandler)

	router.GET(routePrefix+"/health", healthCheck)

}

func apiPadhaiWebhookHandler(c *gin.Context) {

	var webhook models.DockerHubWebhook

	// Parse the JSON payload (optional)
	if err := c.ShouldBindJSON(&webhook); err != nil {
		log.Printf("Warning: Failed to parse webhook JSON: %v", err)
		// Continue anyway since we're restarting a specific container
	} else {
		log.Printf("Received webhook for repository: %s, tag: %s",
			webhook.Repository.Name, webhook.PushData.Tag)
	}

	// Restart the container
	dockerRestartModel := models.DockerContainerModel{
		ContainerName: "uditnair90_api-padhai-golang",
		ImageName:     "uditnair90/api-padhai-golang:latest",
		Port:          "10000",
		ShouldPull:    true,
	}

	if err := controllers.RestartContainer(dockerRestartModel); err != nil {
		log.Printf("Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Failed to restart container: %v", err),
		})
		return
	}

	// Return success
	log.Println("Container restart completed successfully")
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Container restarted successfully",
	})
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "up",
		"message": "Service is running",
	})
}
