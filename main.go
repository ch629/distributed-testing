package main

import (
	"context"
	"distributed-testing/container"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"net/http"
	"os"
	"time"
)

type (
	Config struct {
	}

	Idk struct {
		Config *Config
		Logger *zap.SugaredLogger
	}
)

func main() {
	_ = NewIdk()
	ctx := context.Background()

	options := &container.MakeContainerOptions{
		ImageHost:     "docker.io",
		ImageName:     "mendhak/http-https-echo",
		ContainerName: "echo-container",
		//Cmd: []string{"cat", "/app/main.go"},
	}

	containerId, err := container.RunContainer(options)

	if err != nil {
		panic(err)
	}

	cli, err := client.NewEnvClient()

	if err != nil {
		panic(err)
	}

	if err = container.WaitForContainer(&container.WaitForContainerOptions{
		HealthCheck: healthCheck,
		Interval:    500 * time.Millisecond,
		Retries:     5,
	}); err != nil {
		panic(err)
	}

	sendCall()

	timeout := 10 * time.Second
	if err = cli.ContainerStop(ctx, containerId, &timeout); err != nil {
		panic(err)
	}

	if err = cli.ContainerRemove(ctx, containerId, types.ContainerRemoveOptions{Force: true}); err != nil {
		panic(err)
	}
}

func healthCheck() bool {
	resp, err := http.Get("http://localhost:8080/test")
	if err != nil {
		return false
	}

	return resp.StatusCode == 200
}

func NewIdk() *Idk {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := config.Build()
	return &Idk{
		Config: &Config{},
		Logger: logger.Sugar(),
	}
}

func sendCall() {
	resp, err := http.Get("http://localhost:8080/hello-world")

	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, resp.Body)
}
