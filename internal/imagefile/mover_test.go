package imagefile_test

import "testing"

// The most straightforward way to test this is to just use real files and directories
// Test files made with Gimp, exif data manipulated with exiftool, e.g.:
//     exiftool "-AllDates=1977:08:18 12:00:00" blue.jpg
// blue.jpg:  1977:08:18
// red.jpg:   1976:03:25
// white.jpg: 2006:06:02

// TODO: start here next time (see test/main.go for strategy for path walking to come up with the comparison to expepctedDest)
func TestExecute(t *testing.T) {
	tests := []struct {
		name         string
		expectedDest map[string]string
	}{
		{
			"happy path",
			map[string]string{
				"1976/03": "red.jpg",
				"1977/08": "blue.jpg",
				"2006/06": "white.jpg",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}

// create a temp directory
// copy jpg files from testdata to the temp directory
// temp directory is the source directory
// create a temp directory as the dest directory
// defer cleanup
