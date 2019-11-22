package application

import (
	"io/ioutil"
	"os"
)

func Version(versionFile string) string {
	defaultVersion := "0.0.0"

	if fileExists(versionFile) {
		ver, err := ioutil.ReadFile(versionFile)
		if err != nil {
			return defaultVersion
		}
		return ver
	}
	return defaultVersion
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
