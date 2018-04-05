package getticks

import (
	"fmt"
	"log"
	"time"

	"github.com/thebotguys/golang-bittrex-api/bittrex"

	cli "gopkg.in/urfave/cli.v1"
)

func printCandleStick(candle *bittrex.CandleStick) {
	fmt.Printf("%d\t", candle.High)
	fmt.Printf("%d\t", candle.Close)
	fmt.Printf("%d\t", candle.Low)
	fmt.Printf("%d\t", candle.Volume)
	fmt.Printf("%d\t", candle.BaseVolume)
	fmt.Printf("%s\t", time.Time(candle.Timestamp).String())
}

// ValidateArg nothing to be validated.
func ValidateArg(c *cli.Context) error {
	return nil
}

// Run get market names.
func Run(cctx *cli.Context) error {
	err := bittrex.IsAPIAlive()
	if err != nil {
		return cli.NewExitError(fmt.Errorf("cannot reach bittrex: %v", err), 86)
	}

	candles, err := bittrex.GetTicks("USDT-BTC", "day")
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, candle := range candles {
		printCandleStick(&candle)
	}

	return nil
}
