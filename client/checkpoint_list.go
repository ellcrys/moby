package client // import "github.com/ellcrys/docker/client"

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/ellcrys/docker/api/types"
)

// CheckpointList returns the checkpoints of the given container in the docker host
func (cli *Client) CheckpointList(ctx context.Context, container string, options types.CheckpointListOptions) ([]types.Checkpoint, error) {
	var checkpoints []types.Checkpoint

	query := url.Values{}
	if options.CheckpointDir != "" {
		query.Set("dir", options.CheckpointDir)
	}

	resp, err := cli.get(ctx, "/containers/"+container+"/checkpoints", query, nil)
	if err != nil {
		return checkpoints, wrapResponseError(err, resp, "container", container)
	}

	err = json.NewDecoder(resp.body).Decode(&checkpoints)
	ensureReaderClosed(resp)
	return checkpoints, err
}
