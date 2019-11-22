package cmd

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

var ErrNotFound = errors.New("Not found")

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func DirExists(dirname string) bool {
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func ResolveDirectoryArg(dir string) (string, error) {
	resolved, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}

	if !DirExists(resolved) {
		return "", errors.Wrapf(ErrNotFound, "Directory: %s", resolved)
	}

	return resolved, nil
}
