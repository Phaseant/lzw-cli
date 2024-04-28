package encode

import "github.com/urfave/cli/v2"

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:     "text-path",
		Usage:    "path to file with text",
		Required: true,
		EnvVars:  []string{"TEXT_PATH"},
	},
	&cli.StringFlag{
		Name:    "dictionary-path",
		Usage:   "path to dictionary file. If not set, default one is used",
		EnvVars: []string{"DICTIONARY_PATH"},
	},
}
