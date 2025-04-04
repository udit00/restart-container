package main

import (
	"fmt"
	"udit/restart-container/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const (
	PORT = "20000"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	appServer := gin.Default()

	handlers.RestartDockerHandler(appServer)

	// Start the server
	fmt.Println("Starting webhook server on port " + PORT)
	if err := appServer.Run(":" + PORT); err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}

}

/* SERVICE */
// [Unit]
// Description=Restart Container Golang Service to run the Golang binary
// After=network.target

// [Service]
// ExecStart=/root/API/restart-container-golang/restart-container-golang-binary
// WorkingDirectory=/root/API/restart-container-golang
// Restart=always
// RestartSec=5
// User=root
// Environment=ENV=production

// [Install]
// WantedBy=multi-user.target

/* SERVICE COMMANDS */
/*
sudo systemctl daemon-reexec
sudo systemctl daemon-reload
sudo systemctl enable restart-container-golang
sudo systemctl start restart-container-golang

sudo systemctl status restart-container-golang
journalctl -u restart-container-golang -f

*/
