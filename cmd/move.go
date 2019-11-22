package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

var Move = cli.Command{
	Name:    "move",
	Aliases: []string{"m"},
	Usage:   "Move files to directories sorted by year and month",
	Flags:   cmd.CommonFlags,
	Action:  MoveAction,
}

func MoveAction(c *cli.Context) error {
	fmt.Printf("MoveAction TBD\n")
	return nil
}
