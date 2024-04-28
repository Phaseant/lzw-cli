package decode

import "github.com/urfave/cli/v2"

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:     "encoded-text-path",
		Usage:    "path to encoded text file",
		Required: true,
		EnvVars:  []string{"ENCODED_TEXT_PATH"},
	},
	&cli.StringFlag{
		Name:    "dictionary-path",
		Usage:   "path to dictionary file. If not set, default one is used",
		EnvVars: []string{"DICTIONARY_PATH"},
	},
}
