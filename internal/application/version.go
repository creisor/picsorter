package application

import (
	"io/ioutil"

	"github.com/creisor/picsorter/cmd"
)

func Version(versionFile string) string {
	defaultVersion := "0.0.0"

	if cmd.FileExists(versionFile) {
		ver, err := ioutil.ReadFile(versionFile)
		if err != nil {
			return defaultVersion
		}
		return string(ver)
	}
	return defaultVersion
}
