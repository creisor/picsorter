package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var Copy = cli.Command{
	Name:    "copy",
	Aliases: []string{"c"},
	Usage:   "Copy files to directories sorted by year and month",
	Flags:   CommonFlags,
	Action:  CopyAction,
}

func CopyAction(c *cli.Context) error {
	fmt.Printf("CopyAction TBD\n")
	return nil
}
