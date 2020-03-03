package util

import (
	"fmt"
	"os"
	"path/filepath"
)

// Fo
func SearchUp(filePath string, baseName string) []string {
	var files []string

	dir, _ := filepath.Abs(filePath)

	if filepath.Base(filePath) == baseName {
		dir = filepath.Dir(dir)
	}

	for {
		if dir == "/" {
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
