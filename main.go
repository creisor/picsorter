package main

import (
	"fmt"
	"log"
	"os"

	"github.com/creisor/picsorter/cmd"
	"github.com/creisor/picsorter/internal/application"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/urfave/cli"
)

var testFile = "IMG_0975.jpg"

var (
	Version string
	debug   = false
)

func main() {
	app := cli.NewApp()
	app.Name = "picsorter"
	app.Usage = "Organizes image files in the source directory into a directory structure based on the year and month from the EXIF data in the files"
	app.Flags = []cli.Flag{}
	app.Version = application.Version("./VERSION")

	app.Commands = []cli.Command{
		cmd.Move,
		cmd.Copy,
		cmd.Link,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	/**********************/

	f, err := os.Open(testFile)
	if err != nil {
		log.Fatal(err)
	}

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	dateTime, err := x.DateTime()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("date/time: %+v\n", dateTime)
}
