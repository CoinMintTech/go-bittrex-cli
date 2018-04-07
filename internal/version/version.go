package version

import (
	"runtime"

	"github.com/golang/dep"
	cli "gopkg.in/urfave/cli.v1"
)

var (
	version    = "devel"
	buildDate  string
	commitHash string
)

const versionHelp = "print version information"

// ValidateArg nothing to be validated.
func ValidateArg(c *cli.Context) error {
	return nil
}

// Run print version information.
func Run(ctx *dep.Ctx) error {
	ctx.Out.Printf(`
version     : %s
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
