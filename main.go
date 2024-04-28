package main

import (
	"os"

	"github.com/phaseant/lzw-cli/cmd/decode"
	"github.com/phaseant/lzw-cli/cmd/encode"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	app := &cli.App{
		Usage: "cli to encode (decode) text to (from) lzw",
		Commands: []*cli.Command{
			&encode.Cmd,
			&decode.Cmd,
		},
		Flags:   []cli.Flag{},
		Version: "1.0",
		OnUsageError: func(c *cli.Context, _ error, _ bool) error {
			return cli.ShowAppHelp(c)
		},
		Before: func(ctx *cli.Context) error {
			cfg := zap.NewDevelopmentConfig()
			cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
			cfg.DisableCaller = true
			cfg.DisableStacktrace = true

			zap.ReplaceGlobals(zap.Must(cfg.Build()))
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		zap.S().Error(err)
	}
}
