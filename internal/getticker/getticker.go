package getticker

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/toorop/go-bittrex"

	cli "gopkg.in/urfave/cli.v1"
)

// resolution number of digits after the decimal point.
const (
	resolution = 8

	APIKey    = ""
	APISecret = ""

	flagInterval = "interval"
)

var (
	interval time.Duration
)

// ValidateArg nothing to be validated.
func ValidateArg(c *cli.Context) error {
	interval = c.Duration(flagInterval)

	if interval < 0 {
		return fmt.Errorf("--%s cannot be a negative number: %d", flagInterval, interval)
	}

	return nil
}

// Run get the current tick values for a market.
func Run(cctx *cli.Context) error {
	c := &http.Client{
		Timeout: time.Second * 10,
	}

	btrc := bittrex.NewWithCustomHttpClient(APIKey, APISecret, c)

	switch interval {
	case time.Duration(0):
		innerLoop(btrc)

	default:
		for {
			innerLoop(btrc)
			time.Sleep(interval)
		}

	}

	return nil
}

func innerLoop(btrc *bittrex.Bittrex) {

	now := time.Now().UTC()
	ticker, err := btrc.GetTicker("USDT-BTC")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR %s %s", time.Now().UTC().String(), err.Error())
		return
	}

	printTicker(&now, &ticker)
}

func printTicker(ts *time.Time, ticker *bittrex.Ticker) {
	fmt.Printf("%d\t", ts.UnixNano())
	fmt.Printf("%s\t", ticker.Ask.StringFixedBank(resolution))
	fmt.Printf("%s\t", ticker.Bid.StringFixedBank(resolution))
	fmt.Printf("%s\n", ticker.Last.StringFixedBank(resolution))
}
