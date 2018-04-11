//
// Copyright 2015-2018 Pedro Salgado
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
