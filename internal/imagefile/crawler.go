package imagefile

import "time"

type SomeStructWithMetadataAndFilenames struct {
	Name     string
	DateTime time.Time
}

type Crawler interface {
	Files() (SomeStructWithMetadataAndFilenames, error)
}

type ImageCrawler struct {
	Directory string
	Getter    MetaGetter
}

func NewImageCrawler(dir string, getter MetaGetter) *ImageCrawler {
	return &ImageCrawler{dir, getter}
}

func (c *ImageCrawler) Files() (SomeStructWithMetadataAndFilenames, error) {
	return SomeStructWithMetadataAndFilenames{}, nil
}
