package decode

import (
	"github.com/phaseant/lzw-cli/internal/dictionary"
	"github.com/phaseant/lzw-cli/internal/service/lzw"
	"github.com/phaseant/lzw-cli/pkg/utils"
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

// todo doesnt work

func run(c *cli.Context) error {
	input, err := utils.ReadFile(c.String("encoded-text-path"))
	if err != nil {
		return err
	}

	dictRaw, err := utils.ReadFile(c.String("dictionary-path"))
	if err != nil {
		return err
	}

	dict, err := dictionary.Open(dictRaw)

	service := lzw.New()

	service.Decode(byteToInt(input), dict)

	return nil
}

func byteToInt(b []byte) []int {
	ints := make([]int, len(b))

	for i := 0; i < len(b); i++ {
		ints = append(ints, int(b[i]))
	}

	return ints
}
