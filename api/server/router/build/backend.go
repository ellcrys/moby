package build // import "github.com/ellcrys/docker/api/server/router/build"

import (
	"context"

	"github.com/ellcrys/docker/api/types"
	"github.com/ellcrys/docker/api/types/backend"
)

// Backend abstracts an image builder whose only purpose is to build an image referenced by an imageID.
type Backend interface {
	// Build a Docker image returning the id of the image
	// TODO: make this return a reference instead of string
	Build(context.Context, backend.BuildConfig) (string, error)

	// Prune build cache
	PruneCache(context.Context) (*types.BuildCachePruneReport, error)
}

type experimentalProvider interface {
	HasExperimental() bool
}
