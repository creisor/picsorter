package cmd

import (
	"fmt"

	"github.com/creisor/picsorter/internal/imagefile"
	"github.com/urfave/cli"
)

var (
	Move = cli.Command{
		Name:    "move",
		Aliases: []string{"m"},
		Usage:   "Move files to directories sorted by year and month",
		Flags:   CommonFlags,
		Action:  MoveAction,
	}
)

func MoveAction(c *cli.Context) error {
	sourceDir, err := ResolveDirectoryArg(c.String("source"))
	if err != nil {
		return err
	}
	fmt.Printf("source directory: %s\n", sourceDir)

	destDir := sourceDir
	if c.String("destination") != "" {
		destDir, err = ResolveDirectoryArg(c.String("destination"))
		if err != nil {
			return err
		}
	}
	fmt.Printf("destination directory: %s\n", destDir)

	metaGetter := imagefile.NewDateTimeGetter()
	crawler := imagefile.NewImageCrawler(sourceDir, metaGetter)
	// TODO: next opendev
	//imageFiles, err := crawler.Files()
	//if err != nil {
	//	return err
	//}
	//organizer := imagefile.NewMover(imageFiles)

	fmt.Printf("MoveAction TBD\n")
	return nil
}
