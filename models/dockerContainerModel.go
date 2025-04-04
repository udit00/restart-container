package models

type DockerContainerModel struct {
	ContainerName string `json:"container_name"`
	ImageName     string `json:"image_name"`
	Port          string `json:"port"`
	ShouldPull    bool   `json:"should_pull"`
}
