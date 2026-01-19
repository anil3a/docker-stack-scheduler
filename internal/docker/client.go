package docker

import (
	"context"

	"github.com/docker/docker/client"
)

func New() (*client.Client, error) {
	return client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
}

func Ping() error {
	cli, err := New()
	if err != nil {
		return err
	}
	_, err = cli.Ping(context.Background())
	return err
}
