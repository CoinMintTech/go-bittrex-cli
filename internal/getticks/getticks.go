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

func printCandleStick(candle *bittrex.CandleStick) {
	fmt.Printf("%s\t", candle.High.StringFixedBank(resolution))
	fmt.Printf("%s\t", candle.Close.StringFixedBank(resolution))
	fmt.Printf("%s\t", candle.Low.StringFixedBank(resolution))
	fmt.Printf("%s\t", candle.Volume.StringFixedBank(resolution))
	fmt.Printf("%s\t", candle.BaseVolume.StringFixedBank(resolution))
	fmt.Printf("%s\n", time.Time(candle.Timestamp))
}

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
