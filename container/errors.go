package container

import "fmt"

type (
	imagePullError struct {
		imageName string
		error     error
	}

	containerCreateError struct {
		imageName string
		error     error
	}

	containerStartError struct {
		imageName string
		error     error
	}

	containerWaitError struct {
		containerId string
		error       error
	}
)

func (e imagePullError) Error() string {
	return fmt.Sprintf("Error: Failed to pull image: %v due to %v", e.imageName, e.error)
}

func (e containerCreateError) Error() string {
	return fmt.Sprintf("Error: Failed to create container with image name: %v due to %v", e.imageName, e.error)
}

func (e containerStartError) Error() string {
	return fmt.Sprintf("Error: Failed to start container with image name: %v due to %v", e.imageName, e.error)
}

func (e containerWaitError) Error() string {
	return fmt.Sprintf("Error: Failed to wait for container to exit container with container id: %v due to %v", e.containerId, e.error)
}
