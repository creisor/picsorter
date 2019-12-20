package cmd

import (
	"fmt"

	"github.com/creisor/picsorter/internal/imagefile"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

var (
	// Move is a cli subcommand for running the MoveAction
	Move = cli.Command{
		Name:    "move",
		Aliases: []string{"m"},
		Usage:   "Move files to directories sorted by year and month",
		Flags:   CommonFlags,
		Action:  MoveAction,
	}
)

// MoveAction is a cli Action func for moving files from a source directory to a destination directory
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

	imageFiles, err := crawler.Files()
	if err != nil {
		return err
	}
	if len(imageFiles) < 1 {
		fmt.Printf("No files found in %s\n", sourceDir)
		return nil
	}

	dryRun := c.Bool("dry-run")
	if dryRun {
		fmt.Println("[DRY RUN] commands will not be executed")
	}

	organizer := imagefile.NewMover(imageFiles, sourceDir, destDir)

	err = organizer.Execute(dryRun)
	if err != nil {
		return errors.Wrap(err, "Failed to execute the move")
	}

	return nil
}
