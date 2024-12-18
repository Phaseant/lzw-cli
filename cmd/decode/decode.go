package decode

import (
	"fmt"
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
	if err != nil {
		return err
	}

	service := lzw.New()

	intInput := byteToInt(input)

	decoded := service.Decode(intInput, dict)

	strRep := string(decoded)
	strRep = strings.Replace(strRep, "\\n", "\n", -1)

	zap.S().Infof("Decoded text: %v", strRep)
	zap.S().Infof("Full dictionary: %v", dict.SortedByKeys())

	sizeOfInput := len(input)
	sizeOfOutput := outputInBinary(strRep)

	zap.S().Infof("Size of encoded: %v, decoded: %v, compressed: %.2f%%", sizeOfInput, sizeOfOutput, (1-float64(sizeOfInput)/float64(sizeOfOutput))*100)
	return nil
}

func byteToInt(b []byte) []uint64 {
	var res []uint64

	splitted := strings.Split(string(b), " ")
	for _, i := range splitted {
		if i == "" {
			continue
		}

		num, err := strconv.ParseUint(i, 2, len(i)*8)
		if err != nil {
			panic(err)
		}
		res = append(res, num)
	}

	return res
}

func outputInBinary(output string) int {
	var sb strings.Builder
	for _, val := range output {
		t := fmt.Sprintf("%08b ", val)
		re := strings.TrimLeft(t, "0")
		if re == " " {
			re = "0 "
		}
		sb.WriteString(re)
	}

	return len(sb.String())
}
