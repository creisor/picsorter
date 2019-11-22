package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rwcarlsen/goexif/exif"
)

var testFile = "IMG_0975.jpg"

func main() {
	f, err := os.Open(testFile)
	if err != nil {
		log.Fatal(err)
	}

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	dateTime, _ := x.Get(exif.DateTime) // normally, don't ignore errors!

	s, err := dateTime.StringVal()
	fmt.Printf("date/time: %s\n", s)
}
