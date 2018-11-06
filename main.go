package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "PDF parser"
	app.Usage = "This app do just parse pdf."
	app.Version = "1.0.0"
	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, v",
		Usage: "Echo version",
	}

	app.Action = func(ctx *cli.Context) error {
		fname := ctx.String("filename")
		verbose := ctx.Bool("verbose")

		return nil
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose, V",
			Usage: "If set, echo debug logs.",
		},
		cli.StringFlag{
			Name:  "filename, f",
			Usage: "Find from the list in `PATH`",
		},
	}

	app.Run(os.Args)
}
