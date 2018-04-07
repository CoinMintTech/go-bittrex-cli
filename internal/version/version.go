package version

import (
	"fmt"
	"runtime"

	cli "gopkg.in/urfave/cli.v1"
)

var (
	version    = "devel"
	buildDate  string
	commitHash string
)

// ValidateArg nothing to be validated.
func ValidateArg(c *cli.Context) error {
	return nil
}

// Run print version information.
func Run(cctx *cli.Context) error {
	fmt.Printf(`version     : %s
build date  : %s
git hash    : %s
go version  : %s
go compiler : %s
platform    : %s/%s
`,
		version,
		buildDate,
		commitHash,
		runtime.Version(),
		runtime.Compiler,
		runtime.GOOS,
		runtime.GOARCH,
	)

	return nil
}

// Version returns version number.
func Version() string {
	return version
}
