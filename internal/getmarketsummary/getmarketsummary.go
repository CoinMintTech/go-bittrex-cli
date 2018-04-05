package getmarketsummary

import (
	"fmt"
	"log"
	"time"

	"github.com/thebotguys/golang-bittrex-api/bittrex"

	cli "gopkg.in/urfave/cli.v1"
)

func printSummary(summary *bittrex.MarketSummary) {
	//The name of the market (e.g. BTC-ETH).
	fmt.Printf("%s\t", summary.MarketName)
	// The 24h high for the market.
	fmt.Printf("%d\t", summary.High)
	// The 24h low for the market.
	fmt.Printf("%d\t", summary.Low)
	// The value of the last trade for the market (in base currency).
	fmt.Printf("%d\t", summary.Last)
	// The current highest bid value for the market.
	fmt.Printf("%d\t", summary.Bid)
	// The current lowest ask value for the market.
	fmt.Printf("%d\t", summary.Ask)
	// The 24h volume of the market, in market currency.
	fmt.Printf("%s\t", summary.Volume)
	// The 24h volume for the market, in base currency.
	fmt.Printf("%s\t", summary.BaseVolume)
	// The timestamp of the request.
	fmt.Printf("%s\t", summary.Timestamp)
	// The number of currently open buy orders.
	fmt.Printf("%d\t", summary.OpenBuyOrders)
	// The number of currently open sell orders.
	fmt.Printf("%d\t", summary.OpenSellOrders)
	// The closing price 24h before.
	fmt.Printf("%s\t", summary.PrevDay)
	// The timestamp of the creation of the market.
	fmt.Printf("%s\t\n", summary.Created)
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

	for i := 0; i < 10; i++ {
		summary, err := bittrex.GetMarketSummary("USDT-BTC")
		if err != nil {
			log.Fatal(err.Error())
		}

		printSummary(summary)

		time.Sleep(time.Second)
	}

	return nil
}
