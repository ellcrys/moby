package session // import "github.com/ellcrys/docker/api/server/router/session"

import (
	"context"
	"net/http"

	"github.com/ellcrys/docker/errdefs"
)

func (sr *sessionRouter) startSession(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	err := sr.backend.HandleHTTPRequest(ctx, w, r)
	if err != nil {
		return errdefs.InvalidParameter(err)
	}
	return nil
}
