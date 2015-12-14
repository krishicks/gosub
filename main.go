package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gosub"
	app.Usage = "go dependency submodule automator"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:      "list",
			ShortName: "e",
			Usage:     "list all packages required by the given packages",
			Flags: []cli.Flag{
				cli.StringSliceFlag{
					Name:  "app, a",
					Value: &cli.StringSlice{},
				},
				cli.StringSliceFlag{
					Name:  "test, t",
					Value: &cli.StringSlice{},
				},
			},
			Action: list,
		},
		{
			Name:      "sync",
			ShortName: "s",
			Usage:     "sync packages as submodules (git), or vendored (other)",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "repo, r",
					Value: ".",
				},
				cli.StringFlag{
					Name:  "gopath, g",
					Value: ".",
				},
				cli.StringSliceFlag{
					Name:  "ignore, i",
					Value: &cli.StringSlice{},
					Usage: "Submodule to ignore",
				},
			},
			Action: sync,
		},
		{
			Name:      "vendor",
			ShortName: "v",
			Usage:     "vendor packages as submodules",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "repo, r",
					Value: ".",
				},
				cli.StringFlag{
					Name:  "gopath, g",
					Value: ".",
				},
			},
			Action: vendor,
		},
		{
			Name:      "fix",
			ShortName: "f",
			Usage:     "fix partially constructed submodules",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "repo, r",
					Value: ".",
				},
				cli.StringFlag{
					Name:  "gopath, g",
					Value: ".",
				},
			},
			Action: fix,
		},
	}

	app.Run(os.Args)
}
