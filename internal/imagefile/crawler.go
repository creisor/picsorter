package imagefile

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// Crawler is an interface for crawling files to gather FileInfo
type Crawler interface {
	Files() (FileInfo, error)
}

// ImageCrawler is an implementation of the Crawler interface
type ImageCrawler struct {
	Directory  string
	DateGetter MetaGetter
	files      []FileInfo
}

// NewImageCrawler returns a pointer to an initialized ImageCrawler
func NewImageCrawler(dir string, getter MetaGetter) *ImageCrawler {
	return &ImageCrawler{dir, getter, []FileInfo{}}
}

// Files returns a list of image files
func (c *ImageCrawler) Files() ([]FileInfo, error) {
	var err error

	// TODO: experiment with concurrency: https://godoc.org/github.com/stretchr/powerwalk
	err = filepath.Walk(c.Directory, c.Walker)

	// TODO: does sorting slow us down? if so, maybe make it optional
	sort.Sort(ByDate(c.files))
	return c.files, err
}

// Walker is a filepath WalkFunc (https://golang.org/pkg/path/filepath/#WalkFunc)
func (c *ImageCrawler) Walker(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}

	// instructing gosec to ignore this since we are checking mime type below
	// #nosec G304 (CWE-22): Potential file inclusion via variable
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	mimeType := http.DetectContentType(data)

	if strings.Split(mimeType, "/")[0] == "image" {
		dt, err := c.DateGetter.Get(bytes.NewReader(data))
		if err != nil {
			return err
		}

		// Get() returns an interface{}, so we must assert the type we expect
		// in this case, it's time.Time
		dateTime, ok := dt.(time.Time)
		if !ok {
			return fmt.Errorf("DateTime was an unexpected type: %T (expected time.Time)", dt)
		}
		info := FileInfo{filepath.Base(path), path, dateTime}
		c.files = append(c.files, info)

		return nil
	}

	return nil

}
