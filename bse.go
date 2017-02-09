package main

import (
	"errors"
	"os"

	"github.com/ericdcobb/blob-store-explorer/explorer"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Explore your Blob Store!"

	app.Action = func(c *cli.Context) error {
		path := c.Args().Get(0)
		if path == "" {
			return errors.New("you must supply the blob store directory you wish to explore")
		}
		exploration := explore.Explore(path)
		exploration.Run()
		return nil
	}

	app.Run(os.Args)
}
