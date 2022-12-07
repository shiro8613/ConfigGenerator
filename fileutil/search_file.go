package fileutil

import "path/filepath"

func SearchFiles(path string) ([]string, error) {
	pattern := path + "/*.yml"
	files, err := filepath.Glob(pattern)

	return files, err
}
