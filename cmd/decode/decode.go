package decode

import (
	"github.com/urfave/cli/v2"
)

var Cmd = cli.Command{
	Name:  "decode",
	Usage: "decode from lzw to string",
	Flags: flags,
	OnUsageError: func(c *cli.Context, err error, isSubCommand bool) error {
		return cli.ShowCommandHelp(c, "decode")
	},
	Action: run,
}

func run(c *cli.Context) error {
	return nil
}
