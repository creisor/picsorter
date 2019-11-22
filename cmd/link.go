package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

var Link = cli.Command{
	Name:    "link",
	Aliases: []string{"l"},
	Usage:   "Create directories sorted by year and month, and fill them with symlinks to the real image files",
	Flags:   CommonFlags,
	Action:  LinkAction,
}

func LinkAction(c *cli.Context) error {
	fmt.Printf("LinkAction TBD\n")
	return nil
}
