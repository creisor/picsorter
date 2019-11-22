package cmd

import "github.com/urfave/cli"

var (
	CommonFlags = []cli.Flag{
		cli.StringFlag{
			Name:   "source, s",
			Usage:  "Source directory for files",
			EnvVar: "SOURCE_DIR,SOURCE",
		},
		cli.StringFlag{
			Name:   "destination, d",
			Usage:  "Destination directory for files.  If not specified, it will be the same as the source directory.",
			EnvVar: "DEST_DIR,DEST",
		},
	}
)
