//
// Copyright 2015-2019 Pedro Salgado
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"log"
	"os"

	cli "gopkg.in/urfave/cli.v1"

	"github.com/CoinMintTech/go-bittrex-cli/internal/buylimit"
	"github.com/CoinMintTech/go-bittrex-cli/internal/getbalance"
	"github.com/CoinMintTech/go-bittrex-cli/internal/getmarkets"
	"github.com/CoinMintTech/go-bittrex-cli/internal/getmarketsummary"
	"github.com/CoinMintTech/go-bittrex-cli/internal/getticker"
	"github.com/CoinMintTech/go-bittrex-cli/internal/getticks"
	"github.com/CoinMintTech/go-bittrex-cli/internal/selllimit"
	"github.com/CoinMintTech/go-bittrex-cli/internal/subscribeexchangeupdate"
	"github.com/CoinMintTech/go-bittrex-cli/internal/version"
)

func main() {
	app := cli.NewApp()

	app.Usage = "bittrex-cli is a command-line interface to the Bittrex API."
	app.Version = version.Version()
	app.Commands = []cli.Command{
		{
			Name:      "balance",
			Usage:     "balance <currency>",
			ArgsUsage: "<currency>",
			Before:    getbalance.ValidateArg,
			Action:    getbalance.Run,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "bittrex-key",
					Usage: "Bittrex key (optional, BITTREX_KEY environment variable).",
				},
				cli.StringFlag{
					Name:  "bittrex-secret",
					Usage: "Bittrex secret (optional, BITTREX_SECRET environment variable).",
				},
			},
		},
		{
			Name:      "buy",
			Usage:     "buy <market> <quantity> <rate>",
			ArgsUsage: "<market> <quantity> <rate>",
			Before:    buylimit.ValidateArg,
			Action:    buylimit.Run,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "bittrex-key",
					Usage: "Bittrex key (optional, BITTREX_KEY environment variable).",
				},
				cli.StringFlag{
					Name:  "bittrex-secret",
					Usage: "Bittrex secret (optional, BITTREX_SECRET environment variable).",
				},
			},
		},
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
					Name:      "subscribeexchangeupdate",
					Usage:     "TODO",
					ArgsUsage: "<market>",
					Before:    subscribeexchangeupdate.ValidateArg,
					Action:    subscribeexchangeupdate.Run,
					Flags:     []cli.Flag{},
				},
				{
					Name:      "ticker",
					Usage:     "get the current tick values for a market",
					ArgsUsage: "<market>",
					Before:    getticker.ValidateArg,
					Action:    getticker.Run,
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
			Name:      "sell",
			Usage:     "sell <market> <quantity> <rate>",
			ArgsUsage: "<market> <quantity> <rate>",
			Before:    selllimit.ValidateArg,
			Action:    selllimit.Run,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "bittrex-key",
					Usage: "Bittrex key (optional, BITTREX_KEY environment variable).",
				},
				cli.StringFlag{
					Name:  "bittrex-secret",
					Usage: "Bittrex secret (optional, BITTREX_SECRET environment variable).",
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
