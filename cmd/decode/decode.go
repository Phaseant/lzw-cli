package decode

import (
	"github.com/phaseant/lzw-cli/internal/dictionary"
	"github.com/phaseant/lzw-cli/internal/service/lzw"
	"github.com/phaseant/lzw-cli/pkg/utils"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"strconv"
	"strings"
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

	decoded := service.Decode(byteToInt(input), dict)

	zap.S().Info(string(decoded))

	return nil
}

func byteToInt(b []byte) []int {
	var res []int

	splitted := strings.Split(string(b), " ")
	for _, i := range splitted {
		if i == "" {
			continue
		}
		num, err := strconv.ParseInt(i, 2, 64)
		if err != nil {
			panic(err)
		}
		res = append(res, int(num))
	}

	return res
}
