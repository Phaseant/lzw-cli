package encode

import (
	"github.com/urfave/cli/v2"
)

var Cmd = cli.Command{
	Name:  "encode",
	Usage: "encode from string to lzw",
	Flags: flags,
	OnUsageError: func(c *cli.Context, err error, isSubCommand bool) error {
		return cli.ShowCommandHelp(c, "encode")
	},
	Action: run,
}

func run(c *cli.Context) error {
	return nil
}
