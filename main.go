package main

import (
	"log"
	"os"

	cli "gopkg.in/urfave/cli.v1"

	"github.com/steenzout/go-bittrex-cli/internal/getmarkets"
	"github.com/steenzout/go-bittrex-cli/internal/getmarketsummary"
	"github.com/steenzout/go-bittrex-cli/internal/getticker"
	"github.com/steenzout/go-bittrex-cli/internal/getticks"
	"github.com/steenzout/go-bittrex-cli/internal/version"
)

func main() {
	app := cli.NewApp()

	app.Usage = "bittrex-cli is a command-line interface to the Bittrex API."
	app.Version = version.Version()
	app.Commands = []cli.Command{
		{
			Name:  "get",
			Usage: "get data",
			Subcommands: []cli.Command{
				{
					Name:   "markets",
					Usage:  "get the open and available trading markets at Bittrex along with other meta data.",
					Before: getmarkets.ValidateArg,
					Action: getmarkets.Run,
				},
				{
					Name:   "marketsummary",
					Usage:  "get the last 24 hour summary of all active exchanges",
					Before: getmarketsummary.ValidateArg,
					Action: getmarketsummary.Run,
				},
				{
					Name:   "ticker",
					Usage:  "get the current tick values for a market",
					Before: getticker.ValidateArg,
					Action: getticker.Run,
					Flags: []cli.Flag{
						cli.DurationFlag{
							Name:  "interval, i",
							Usage: "the polling interval",
						},
					},
				},
				{
					Name:   "ticks",
					Usage:  "get ticks",
					Before: getticks.ValidateArg,
					Action: getticks.Run,
				},
			},
		},
		{
			Name:   "version",
			Usage:  "version",
			Before: version.ValidateArg,
			Action: version.Run,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
