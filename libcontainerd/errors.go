package libcontainerd // import "github.com/ellcrys/docker/libcontainerd"

import (
	"errors"

	"github.com/ellcrys/docker/errdefs"
)

func newNotFoundError(err string) error { return errdefs.NotFound(errors.New(err)) }

func newInvalidParameterError(err string) error { return errdefs.InvalidParameter(errors.New(err)) }

func newConflictError(err string) error { return errdefs.Conflict(errors.New(err)) }
