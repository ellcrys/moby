package config // import "github.com/ellcrys/docker/daemon/config"

import (
	"github.com/ellcrys/docker/api/types/swarm"
	"github.com/ellcrys/docker/daemon/cluster/convert"
	"github.com/docker/swarmkit/api/genericresource"
)

// ParseGenericResources parses and validates the specified string as a list of GenericResource
func ParseGenericResources(value []string) ([]swarm.GenericResource, error) {
	if len(value) == 0 {
		return nil, nil
	}

	resources, err := genericresource.Parse(value)
	if err != nil {
		return nil, err
	}

	obj := convert.GenericResourcesFromGRPC(resources)
	return obj, nil
}
