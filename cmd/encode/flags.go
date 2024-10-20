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
		Name:     "dictionary-path",
		Usage:    "path to dictionary file",
		Required: true,
		EnvVars:  []string{"DICTIONARY_PATH"},
	},
	&cli.StringFlag{
		Name:    "output-path",
		Usage:   "path to output file to be saved",
		EnvVars: []string{"OUTPUT_PATH"},
	},
}
