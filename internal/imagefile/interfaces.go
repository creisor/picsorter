package imagefile

import "io"

// MetaGetter is an interface for getting metadata from a file
type MetaGetter interface {
	Get(io.Reader) (interface{}, error)
}

// Organizer is an interface for organizing files
type Organizer interface {
	Execute(bool) error
}
