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

package subscribeexchangeupdate

import (
	"fmt"

	"github.com/toorop/go-bittrex"

	cli "gopkg.in/urfave/cli.v1"
)

// resolution number of digits after the decimal point.
const resolution = 8

// ValidateArg nothing to be validated.
func ValidateArg(c *cli.Context) error {
	return nil
}

// Run get market names.
func Run(cctx *cli.Context) error {
	api := bittrex.New("", "")

	dataChan := make(chan bittrex.ExchangeState, 10)
	exitChan := make(chan bool)
	api.SubscribeExchangeUpdate("USDT-BTC", dataChan, exitChan)

	for {
		select {
		case <-exitChan:
			return nil

		default:
			select {
			case <-exitChan:
				return nil
			}

		case update := <-dataChan:
			fmt.Printf("%+v\n", update)
		}
	}
}
