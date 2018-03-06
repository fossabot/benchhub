package runner

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/dyweb/gommon/errors"

	"github.com/benchhub/benchhub/pkg/common/spec"
)

type Docker struct {
	spec spec.Docker
	c    *client.Client
}

func NewDocker(s spec.Docker, c *client.Client) (*Docker, error) {
	var err error
	if c == nil {
		if c, err = client.NewEnvClient(); err != nil {
			return nil, errors.Wrap(err, "can't create docker client from environment")
		}
	}
	return &Docker{
		spec: s,
		c:    c,
	}, nil
}

func (d *Docker) Run(ctx context.Context) error {
	switch d.spec.Action {
	case spec.DockerPull:
		return d.Pull(ctx)
	case spec.DockerRun:
		return d.Start(ctx)
	}
	return errors.New("unknown docker action")
}

// Start starts the container in background, like docker run -d
func (d *Docker) Start(ctx context.Context) error {
	res, err := d.c.ContainerCreate(ctx, &container.Config{
		Image: d.spec.Image,
		// TODO: cmd, tty etc.
	}, nil, nil, "")
	if err != nil {
		return errors.Wrap(err, "failed to create container")
	}
	if err := d.c.ContainerStart(ctx, res.ID, types.ContainerStartOptions{}); err != nil {
		return errors.Wrap(err, "failed to start container")
	}
	// TODO: keep the container id somewhere
	return nil
}

func (d *Docker) Pull(ctx context.Context) error {
	res, err := d.c.ImagePull(ctx, d.spec.Image, types.ImagePullOptions{})
	if err != nil {
		return errors.Wrapf(err, "failed to pull image %s", d.spec.Image)
	}
	// TODO: should copy the output to somewhere instead of just stdout ...
	// NOTE: this is blocking, it will stream JSON from server about status
	// TODO: what happens for a non existing image
	if _, err := io.Copy(os.Stdout, res); err != nil {
		return errors.Wrap(err, "failed to read pull image response")
	}
	return nil
}