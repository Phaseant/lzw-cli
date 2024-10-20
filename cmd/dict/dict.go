package dict

import (
	"github.com/phaseant/lzw-cli/internal/dictionary"
	"github.com/phaseant/lzw-cli/pkg/utils"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var Cmd = cli.Command{
	Name:  "generate-dict",
	Usage: "generate base dictionary",
	Flags: flags,
	OnUsageError: func(c *cli.Context, err error, isSubCommand bool) error {
		return cli.ShowCommandHelp(c, "generate-dict")
	},
	Action: run,
}

func run(c *cli.Context) error {
	input, err := utils.ReadFile(c.String("text-path"))
	if err != nil {
		return err
	}

	d, err := dictionary.New(input)
	if err != nil {
		return err
	}

	raw, err := d.Marshall()
	if err != nil {
		return err
	}

	zap.S().Info("Your dictionary\n", string(raw))

	if err := utils.WriteFile(c.String("dictionary-path"), raw); err != nil {
		return err
	}

	return nil
}
