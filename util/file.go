package util

import (
	"fmt"
	"path/filepath"
)

func SearchUp(filename string) []string {
	var files []string
	fileName := "terradude.hcl"
	dir := filename

	if filepath.Base(filename) == fileName {
		dir = filepath.Dir(dir)
	}

	for {
		if dir == ".git" {
			files = append(files, fmt.Sprintf("%v", fileName))
			break
		}
		files = append(files, fmt.Sprintf("%v/%v", dir, fileName))
		dir = filepath.Dir(dir)
	}
	return files
}
