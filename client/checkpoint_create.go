package client

import (
	"time"

	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)

// CheckpointCreate creates a checkpoint from the given container with the given name
func (cli *Client) CheckpointCreate(ctx context.Context, container string, options types.CheckpointCreateOptions) error {
	//matt's modification --start
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	fmt.Println("docker_03 time is ", tm.Format("2006-01-02 03:04:05:55 PM"))
	//matt's modification --end

	resp, err := cli.post(ctx, "/containers/"+container+"/checkpoints", nil, options, nil)

	//matt's modification --start
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	fmt.Println("docker_04 time is ", tm.Format("2006-01-02 03:04:05:55 PM"))
	//matt's modification --end
	ensureReaderClosed(resp)
	return err
}
