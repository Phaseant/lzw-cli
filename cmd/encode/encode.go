package encode

import (
	"fmt"
	"github.com/phaseant/lzw-cli/internal/dictionary"
	"github.com/phaseant/lzw-cli/internal/service/lzw"
	"github.com/phaseant/lzw-cli/pkg/utils"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"strings"
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
	input, err := utils.ReadFile(c.String("text-path"))
	if err != nil {
		return err
	}

	rawDict, err := utils.ReadFile(c.String("dictionary-path"))

	dict, err := dictionary.Open(rawDict)
	if err != nil {
		return err
	}

	service := lzw.New()

	output := service.Encode(input, dict)

	var sb strings.Builder
	if outputPath := c.String("output-path"); outputPath != "" {
		for _, val := range output {
			t := fmt.Sprintf("%08b ", val)
			re := strings.TrimLeft(t, "0")
			if re == " " {
				re = "0 "
			}
			sb.WriteString(re)
		}

		if err := utils.WriteFile(outputPath, []byte(sb.String())); err != nil {
			return err
		}
	}

	zap.S().Infof("Encoded text: %v", output)
	return nil
}
