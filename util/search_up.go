package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func SearchUp(filePath string, baseName string) []string {
	var files []string

	dir := filePath

	if filepath.Base(filePath) == baseName {
		dir = filepath.Dir(dir)
	}

	for {
		if dir == "." || dir == "/" {
			if _, err := os.Stat(baseName); err == nil {
				return append(files, baseName)
			}
			break
		}
		current := fmt.Sprintf("%v/%v", dir, baseName)
		if _, err := os.Stat(current); err == nil {
			files = append(files, current)
		}

		dir = filepath.Dir(dir)
	}
	return files
}
