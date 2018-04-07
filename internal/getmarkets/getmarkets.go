package getmarkets

import (
	"fmt"

	"github.com/thebotguys/golang-bittrex-api/bittrex"

	cli "gopkg.in/urfave/cli.v1"
)

// ValidateArg nothing to be validated.
func ValidateArg(c *cli.Context) error {
	return nil
}

// Run get market names.
func Run(cctx *cli.Context) error {
	markets, err := bittrex.GetMarkets()
	if err != nil {
		return cli.NewExitError(fmt.Errorf("failed to get markets: %v", err), 86)
	}

	for _, m := range markets {
		fmt.Println(m.MarketName)
	}

	return nil
}
