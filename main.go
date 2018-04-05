package main

import (
	"log"
	"os"

	"github.com/steenzout/go-agent-bittrex-firehose/internal/getmarkets"
	"github.com/steenzout/go-agent-bittrex-firehose/internal/getmarketsummary"
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
					Usage:  "get markets",
					Before: getmarkets.ValidateArg,
					Action: getmarkets.Run,
				},
				{
					Name:   "marketsummary",
					Usage:  "get market summary",
					Before: getmarketsummary.ValidateArg,
					Action: getmarketsummary.Run,
				},
				{
					Name:   "ticks",
					Usage:  "get ticks",
					Before: getticks.ValidateArg,
					Action: getticks.Run,
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
