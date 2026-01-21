package docker

import (
	"context"

	"github.com/docker/docker/api/types"
)

func ListContainers(all bool) ([]types.Container, error) {
	cli, err := New()
	if err != nil {
		return nil, err
	}

	return cli.ContainerList(context.Background(), types.ContainerListOptions{All: all})
}
