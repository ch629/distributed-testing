package main

import (
	"distributed-testing/scenario/scanner"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"net/http"
	"os"
	"path"
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
	wd, _ := os.Getwd()
	file, err := os.Open(path.Join(wd, "example.feature"))

	if err != nil {
		panic(err)
	}

	scanner := scanner.NewScanner(file)
	for {
		if step := scanner.Scan(); step != nil {
			fmt.Println("Step: ", *step)
			continue
		}
		break
	}

	file.Close()
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

	if _, err = io.Copy(os.Stdout, resp.Body); err != nil {
		panic(err)
	}
}
