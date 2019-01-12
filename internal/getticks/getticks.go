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

package getticks

import (
	"fmt"
	"log"
	"time"

	"github.com/thebotguys/golang-bittrex-api/bittrex"

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
	cs, err := bittrex.GetTicks("USDT-BTC", "day")
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, candle := range cs {
		printCandleStick(&candle)
	}

	return nil
}

func printCandleStick(candle *bittrex.CandleStick) {
	fmt.Printf("%s\t", candle.High.StringFixedBank(resolution))
	fmt.Printf("%s\t", candle.Close.StringFixedBank(resolution))
	fmt.Printf("%s\t", candle.Low.StringFixedBank(resolution))
	fmt.Printf("%s\t", candle.Volume.StringFixedBank(resolution))
	fmt.Printf("%s\t", candle.BaseVolume.StringFixedBank(resolution))
	fmt.Printf("%s\n", time.Time(candle.Timestamp))
}
