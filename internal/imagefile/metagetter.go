package imagefile

import (
	"io"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

// DateTimeGetter is an implementation of the MetaGetter interface
type DateTimeGetter struct{}

// NewDateTimeGetter returns a pointer to an instantiated DateTimeGetter
func NewDateTimeGetter() *DateTimeGetter {
	return &DateTimeGetter{}
}

// Get returns DateTime metadata as time.Time
func (g *DateTimeGetter) Get(data io.Reader) (interface{}, error) {
	x, err := exif.Decode(data)
	if err != nil {
		return time.Time{}, err
	}

	d, err := x.DateTime()
	if err != nil {
		return time.Time{}, err
	}

	return d, nil
}
