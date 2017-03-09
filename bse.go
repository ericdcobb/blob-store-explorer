package main

import (
	"os"

	"github.com/ericdcobb/blob-store-explorer/explorer"
	"github.com/urfave/cli"
)

func main() {
	var collect bool
	var format string
	var before string
	var after string

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
		cli.StringFlag{
			Name:        "format, fmt",
			Value:       "text",
			Usage:       "Format of output, 'json' or 'text'",
			Destination: &format,
		},
		cli.StringFlag{
			Name:        "before, b",
			Value:       "",
			Usage:       "Filter to include blobs created before this date",
			Destination: &before,
		},
		cli.StringFlag{
			Name:        "after, a",
			Value:       "",
			Usage:       "Filter to include blobs created after this date",
			Destination: &after,
		},
	}

	app.Action = func(c *cli.Context) error {
		path := ""
		if c.NArg() > 0 {
			path = c.Args()[0]
		}

		exploration := explore.Explore(path, collect, c.StringSlice("filter"), format,
			before, after)
		exploration.Run()
		return nil
	}

	app.Run(os.Args)
}
