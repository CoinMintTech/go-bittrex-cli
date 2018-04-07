package main

import (
	"log"
	"os"

	"github.com/steenzout/go-agent-bittrex-firehose/internal/getmarkets"
	"github.com/steenzout/go-agent-bittrex-firehose/internal/getmarketsummary"
	"github.com/steenzout/go-agent-bittrex-firehose/internal/getticker"
	"github.com/steenzout/go-agent-bittrex-firehose/internal/getticks"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()

	app.Usage = "bittrex-cli is a command-line interface to the Bittrex API."
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
			{
				Name:   "version",
				Usage:  "version",
				Before: getticks.ValidateArg,
				Action: getticks.Run,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
