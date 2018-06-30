package util

import (
	"github.com/docker/docker/client"
)

func Client() (cli *client.Client) {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.37"))
	if err != nil {
		panic(err)
	}
	return
}
