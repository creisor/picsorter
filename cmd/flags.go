package cmd

import "github.com/urfave/cli/v2"

var (
	CommonFlags = []cli.Flag{
		&cli.StringFlag{
			Name:     "source, s",
			Usage:    "Source directory for files",
			EnvVars:  []string{"SOURCE_DIR,SOURCE"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "destination, d",
			Usage:    "Destination directory for files.  If not specified, it will be the same as the source directory.",
			EnvVars:  []string{"DEST_DIR,DEST"},
			Required: true,
		},
		&cli.BoolFlag{
			Name:    "dry-run",
			Aliases: []string{"D"},
			Usage:   "Do not execute, print what would have been done",
			EnvVars: []string{"DRY_RUN"},
			Value:   false,
		},
	}
)
