package container

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"io"
	"os"
	"time"
)

type (
	MakeContainerOptions struct {
		ImageHost     string
		ImageName     string
		ImageTag      string
		ContainerName string
		Cmd           []string
	}
)

// TODO: Validate the / between host & name
func makeImageUrl(imageHost string, imageName string, imageTag string) string {
	if imageHost == "" {
		imageHost = "docker.io/library"
	}

	if imageTag == "" {
		imageTag = "latest"
	}

	return fmt.Sprintf("%v/%v:%v", imageHost, imageName, imageTag)
}

func pullImage(cli *client.Client, image string) error {
	reader, err := cli.ImagePull(context.Background(), image, types.ImagePullOptions{})

	if err != nil {
		return imagePullError{
			imageName: image,
			error:     err,
		}
	}

	defer reader.Close()

	// TODO: Where do we want the reader to go?
	if _, err = io.Copy(os.Stdout, reader); err != nil {
		return err
	}

	return nil
}

// TODO: Build image from a Dockerfile
func buildImage(cli *client.Client) error {
	_, err := cli.ImageBuild(context.Background(), nil, types.ImageBuildOptions{})

	if err != nil {
		return err
	}

	return nil
}

func createContainer(cli *client.Client, imageName string, cmd []string, containerName string) (resp container.ContainerCreateCreatedBody, err error) {
	dir, err := os.Getwd()
	if err != nil {
		return
	}

	containerPort, err := nat.NewPort("tcp", "80")

	if err != nil {
		return
	}

	resp, err = cli.ContainerCreate(context.Background(), &container.Config{
		Image: imageName,
		Cmd:   cmd,
		ExposedPorts: nat.PortSet{
			containerPort: {},
		},
	}, &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: dir + "/",
				Target: "/bind",
			},
		},
		PortBindings: nat.PortMap{
			containerPort: []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "8080",
				},
			},
		},
	}, nil, containerName)

	if err != nil {
		err = containerCreateError{
			imageName: imageName,
			error:     err,
		}
	}
	return
}

func RunContainer(options *MakeContainerOptions) (string, error) {
	ctx := context.Background()

	cli, err := client.NewEnvClient()

	if err != nil {
		return "", err
	}

	if err = pullImage(cli, makeImageUrl(options.ImageHost, options.ImageName, options.ImageTag)); err != nil {
		return "", err
	}

	resp, err := createContainer(cli, options.ImageName, options.Cmd, options.ContainerName)

	if err != nil {
		return "", err
	}

	//defer cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{Force: true})

	if len(resp.Warnings) > 0 {
		for _, warning := range resp.Warnings {
			fmt.Printf("Warning: %v", warning)
		}
	}

	// TODO: We need to see any errors from the logs when starting
	if err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return "", containerStartError{
			imageName: options.ImageName,
			error:     err,
		}
	}

	//if err = AttachContainerStdout(cli, resp.ID); err != nil {
	//	panic(err)
	//}
	//
	//if err = WaitForContainerToExit(cli, resp.ID); err != nil {
	//	return "", err
	//}

	return resp.ID, nil
}

func AttachContainer(cli *client.Client, containerId string) (io.Reader, error) {
	attachResp, err := cli.ContainerAttach(context.Background(), containerId, types.ContainerAttachOptions{
		Stream: true,
		Stdout: true,
		Stderr: true,
		Logs:   true,
	})

	if err != nil {
		return nil, err
	}

	return attachResp.Reader, nil
}

func AttachContainerStdout(cli *client.Client, containerId string) error {
	attachResp, err := cli.ContainerAttach(context.Background(), containerId, types.ContainerAttachOptions{
		Stream: true,
		Stdout: true,
		Stderr: true,
		Logs:   true,
	})

	if err != nil {
		return err
	}

	// TODO: When to close this?
	defer attachResp.Close()
	_, err = io.Copy(os.Stdout, attachResp.Reader)
	return err
}

func StopContainer(cli *client.Client, containerId string) error {
	timeout := 10 * time.Second
	return cli.ContainerStop(context.Background(), containerId, &timeout)
}

func WaitForContainerToExit(cli *client.Client, containerId string) error {
	if _, err := cli.ContainerWait(context.Background(), containerId); err != nil {
		return containerWaitError{
			containerId: containerId,
			error:       err,
		}
	}

	return nil
}

func GetContainerLogs(cli *client.Client, containerId string) (io.ReadCloser, error) {
	return cli.ContainerLogs(context.Background(), containerId, types.ContainerLogsOptions{ShowStdout: true})
}
