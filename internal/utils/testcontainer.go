package utils

import (
	"context"
	"fmt"
	"io"
	"time"

	. "github.com/onsi/ginkgo/v2"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func CreateContainer(ctx context.Context, configFile string, networks []string) (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Image:        "quay.io/oauth2-proxy/oauth2-proxy:latest",
		ExposedPorts: []string{"4180:4180"},
		Hostname:     "oauth2-proxy",
		WaitingFor:   wait.ForListeningPort("4180").WithStartupTimeout(30 * time.Second),
		Files: []testcontainers.ContainerFile{
			{
				HostFilePath:      configFile,
				ContainerFilePath: "/oauth2-proxy.cfg",
				FileMode:          0644,
			},
		},
		Cmd:      []string{"--config", "/oauth2-proxy.cfg"},
		Networks: networks,
	}

	c, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		logs, err := c.Logs(ctx)
		if err == nil {
			defer logs.Close()
			GinkgoWriter.Println("\n=== Container Logs ===")
			_, _ = io.Copy(GinkgoWriter, logs)
			GinkgoWriter.Println("=====================")
		}
		return c, fmt.Errorf("Failed to start container: %v", err)
	}

	time.Sleep(time.Millisecond * 250)
	return c, nil
}
