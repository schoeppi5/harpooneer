package docker

import (
	"github.com/schoeppi5/harpooneer/logging"
	"github.com/schoeppi5/harpooneer/plugin"
)

type DockerExtension struct{}

func NewDockerExtension(log logging.Logger) plugin.Extension {
	return &DockerExtension{}
}

func (de DockerExtension) Name() string {
	return "whatever"
}
