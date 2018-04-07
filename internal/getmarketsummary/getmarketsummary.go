package getmarketsummary

import (
	"fmt"
	"log"
	"time"

	"github.com/thebotguys/golang-bittrex-api/bittrex"

	cli "gopkg.in/urfave/cli.v1"
)

// resolution number of digits after the decimal point.
const resolution = 8

var market string = "USDT-BTC"

func printSummary(summary *bittrex.MarketSummary) {
	//The name of the market (e.g. BTC-ETH).
	fmt.Printf("%s\t", summary.MarketName)
	// The 24h high for the market.
	fmt.Printf("%s\t", summary.High.StringFixedBank(resolution))
	// The 24h low for the market.
	fmt.Printf("%s\t", summary.Low.StringFixedBank(resolution))
	// The value of the last trade for the market (in base currency).
	fmt.Printf("%s\t", summary.Last.StringFixedBank(resolution))
	// The current highest bid value for the market.
	fmt.Printf("%s\t", summary.Bid.StringFixedBank(resolution))
	// The current lowest ask value for the market.
	fmt.Printf("%s\t", summary.Ask.StringFixedBank(resolution))
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
	fmt.Printf("%s\t", summary.PrevDay.StringFixedBank(resolution))
	// The timestamp of the creation of the market.
	fmt.Printf("%s\n", summary.Created)
}

// ValidateArg nothing to be validated.
func ValidateArg(c *cli.Context) error {
	return nil
}

// Run get the last 24 hour summary of all active exchanges.
func Run(cctx *cli.Context) error {
	for i := 0; i < 10; i++ {
		summary, err := bittrex.GetMarketSummary(market)
		if err != nil {
			log.Fatal(err.Error())
		}

		if summary.MarketName != market {
			// bug on v2 API may lead to getting the wrong summary
			continue
		}

		printSummary(summary)

		time.Sleep(2 * time.Second)
	}

	return nil
}
