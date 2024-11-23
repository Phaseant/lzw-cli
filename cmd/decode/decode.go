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

	intInput := byteToInt(input)

	decoded := service.Decode(intInput, dict)

	strRep := string(decoded)
	strRep = strings.Replace(strRep, "\\n", "\n", -1)

	zap.S().Infof("Decoded text: %v", strRep)

	sizeOfInput := len(intInput)
	sizeOfOutput := len(strRep)

	zap.S().Infof("Size of encoded: %v, decoded: %v, compressed: %.2f%%", sizeOfInput, sizeOfOutput, float64(sizeOfInput)/float64(sizeOfOutput)*100)
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
