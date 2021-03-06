package main

import (
	"fmt"
	"log"
	"os"

	"github.com/creisor/picsorter/cmd"

	"github.com/urfave/cli/v2"
)

var (
	Version string
	Build   string
	debug   = false
)

func main() {
	app := cli.NewApp()
	app.Name = "picsorter"
	app.Usage = "Organizes image files in the source directory into a directory structure based on the year and month from the EXIF data in the files"
	app.Flags = []cli.Flag{}
	app.Version = Version

	cli.AppHelpTemplate = fmt.Sprintf(`%s

  BUILD: %s

  `, cli.AppHelpTemplate, Build)

	app.Commands = []*cli.Command{
		&cmd.Move,
		&cmd.Copy,
		&cmd.Link,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
