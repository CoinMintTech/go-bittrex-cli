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

package getbalance

import (
	"fmt"

	"github.com/steenzout/go-env"
	"github.com/toorop/go-bittrex"

	cli "gopkg.in/urfave/cli.v1"
)

const (
	// resolution number of digits after the decimal point.
	resolution = 8

	// flags
	flagKey    = "key"
	flagSecret = "secret"

	// EnvBittrexKey environment variable containing the Bittrex key.
	EnvBittrexKey = "BITTREX_KEY"
	// EnvBittrexSecret environment variable containing the Bittrex secret.
	EnvBittrexSecret = "BITTREX_SECRET"
)

var (
	// arguments
	currency string
	key      string
	secret   string
)

// Key returns the Bittrex jey.
func Key(c *cli.Context) string {
	v := c.String(flagKey)
	if v == "" {
		return env.GetString(EnvBittrexKey)
	}

	return v
}

// Secret returns the Bittrex secret.
func Secret(c *cli.Context) string {
	v := c.String(flagSecret)
	if v == "" {
		return env.GetString(EnvBittrexSecret)
	}

	return v
}

// ValidateArg validate input arguments
func ValidateArg(c *cli.Context) error {
	if c.NArg() != 1 {
		return fmt.Errorf("invalid number of arguments")
	}

	currency = c.Args().First()

	secret = Secret(c)

	key = Key(c)

	return nil
}

// Run request a BUY order.
func Run(cctx *cli.Context) error {
	api := bittrex.New(key, secret)

	balance, err := api.GetBalance(currency)
	if err != nil {
		return err
	}

	fmt.Printf(
		"Available: %s\nBalance: %s\nCryptoAddress: %s\nCurrency: %s\nPending: %s\nRequested: %s\nUUID: %s\n",
		balance.Available.StringFixed(resolution),
		balance.Balance.StringFixed(resolution),
		balance.CryptoAddress,
		balance.Currency,
		balance.Pending.StringFixed(resolution),
		balance.Requested,
		balance.Uuid,
	)

	return nil
}
