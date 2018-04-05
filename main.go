package main

import (
	"fmt"
	"log"
	"time"

	"github.com/thebotguys/golang-bittrex-api/bittrex"
)

func printCandleStick(candle *bittrex.CandleStick) {
	fmt.Printf("High=%s\n", candle.High.StringFixed(8))
	fmt.Printf("Close=%s\n", candle.Close.StringFixed(8))
	fmt.Printf("Low=%s\n", candle.Low.StringFixed(8))
	fmt.Printf("Volume=%s\n", candle.Volume.StringFixed(8))
	fmt.Printf("BaseVolume=%s\n", candle.BaseVolume.StringFixed(8))
	fmt.Printf("Timestamp=%s\n", time.Time(candle.Timestamp).String())
}

func printSummary(summary *bittrex.MarketSummary) {
	//The name of the market (e.g. BTC-ETH).
	fmt.Printf("MarketName=%s\n", summary.MarketName)
	// The 24h high for the market.
	fmt.Printf("High=%s\n", summary.High.StringFixed(8))
	// The 24h low for the market.
	fmt.Printf("Low=%s\n", summary.Low.StringFixed(8))
	// The value of the last trade for the market (in base currency).
	fmt.Printf("Last=%s\n", summary.Last.StringFixed(8))
	// The current highest bid value for the market.
	fmt.Printf("Bid=%s\n", summary.Bid.StringFixed(8))
	// The current lowest ask value for the market.
	fmt.Printf("Ask=%s\n", summary.Ask.StringFixed(8))
	// The 24h volume of the market, in market currency.
	fmt.Printf("Volume=%s\n", summary.Volume)
	// The 24h volume for the market, in base currency.
	fmt.Printf("BaseVolume=%s\n", summary.BaseVolume)
	// The timestamp of the request.
	fmt.Printf("Timestamp=%s\n", summary.Timestamp)
	// The number of currently open buy orders.
	fmt.Printf("OpenBuyOrders=%d\n", summary.OpenBuyOrders)
	// The number of currently open sell orders.
	fmt.Printf("OpenSellOrders=%d\n", summary.OpenSellOrders)
	// The closing price 24h before.
	fmt.Printf("PrevDay=%s\n", summary.PrevDay)
	// The timestamp of the creation of the market.
	fmt.Printf("Created=%s\n", summary.Created)
}

func main() {
	err := bittrex.IsAPIAlive()
	if err != nil {
		fmt.Println("CANNOT REACH BITTREX API SERVERS: ", err)
	}

	markets, err := bittrex.GetMarkets()
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, m := range markets {
		fmt.Printf("market name: %s\n", m.MarketName)
	}

	fmt.Println("")
	fmt.Println("-- GetMarketSummary --")
	fmt.Println("")

	for i := 0; i < 10; i++ {
		summary, err := bittrex.GetMarketSummary("USDT-BTC")
		if err != nil {
			log.Fatal(err.Error())
		}

		printSummary(summary)
		fmt.Println()

		time.Sleep(time.Second)
	}

	fmt.Println("")
	fmt.Println("-- GetTicks --")
	fmt.Println("")

	candles, err := bittrex.GetTicks("USDT-BTC", "day")
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, candle := range candles {
		printCandleStick(&candle)
		fmt.Println("")
	}
}
