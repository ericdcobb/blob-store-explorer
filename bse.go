package main

import (
	"fmt"
	"os"

	"github.com/ericdcobb/blob-store-explorer/explorer"
	"github.com/urfave/cli"
)

func main() {
	var collect bool
	app := cli.NewApp()
	app.Usage = "Explore your Blob Store!"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "collect, c",
			Usage:       "Collect and print the results of the exploration",
			Destination: &collect,
		},
		cli.StringSliceFlag{
			Name: "filter, f",
		},
	}

	app.Action = func(c *cli.Context) error {
		path := ""
		if c.NArg() > 0 {
			path = c.Args()[0]
		}
		fmt.Println(c.StringSlice("filter"))
		exploration := explore.Explore(path, collect, c.StringSlice("filter"))
		exploration.Run()
		return nil
	}

	app.Run(os.Args)
}
