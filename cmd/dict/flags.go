package dict

import "github.com/urfave/cli/v2"

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:     "text-path",
		Usage:    "path to text file",
		Required: true,
		EnvVars:  []string{"ENCODED_TEXT_PATH"},
	},
	&cli.StringFlag{
		Name:     "dictionary-path",
		Usage:    "path to dictionary file.",
		Required: true,
		EnvVars:  []string{"DICTIONARY_PATH"},
	},
}
