package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// DELETEME
// this is me playing around with Walk() and filepath to figure out how to generate the map I want to use for comparison in mover_test.go
// it can be removed later when I'm done writing the tests

// The output of this was:
//
//fileStructure: map[2012/03:IMG_0184.jpg 2012/09:IMG_0462.JPG 2012/10:IMG_0807.jpg 2016/10:SHOULD_BE_2016_10_15.jpg 2017/06:IMG_0975.jpg 2019/05:IMG_0083.JPG 2019/08:IMG_1879.jpg 2019/10:IMG_2945.jpg 2019/11:IMG_3129.jpg]
//2012/03: IMG_0184.jpg
//2012/09: IMG_0462.JPG
//2016/10: SHOULD_BE_2016_10_15.jpg
//2017/06: IMG_0975.jpg
//2019/08: IMG_1879.jpg
//2019/10: IMG_2945.jpg
//2019/11: IMG_3129.jpg
//2012/10: IMG_0807.jpg
//2019/05: IMG_0083.JPG

func main() {
	fileStructure := make(map[string]string)
	baseDir := "/Users/creisor/tmp/picsorter"

	if err := os.Chdir(baseDir); err != nil {
		log.Fatal(err)
	}

	err := filepath.Walk("/Users/creisor/tmp/picsorter", func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}
		dirname, err := filepath.Rel("/Users/creisor/tmp/picsorter", filepath.Dir(path))
		if err != nil {
			return err
		}
		fname := filepath.Base(path)

		fileStructure[dirname] = fname

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("fileStructure: %+v\n", fileStructure)
	for k, v := range fileStructure {
		fmt.Printf("%s: %s\n", k, v)
	}
}
