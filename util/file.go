package util

import (
	"fmt"
	"io/ioutil"
	"os"
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
		if dir == "." {
			break
		}
		current := fmt.Sprintf("%v/%v", dir, fileName)
		if _, err := os.Stat(current); err == nil {
			files = append(files, current)
		}

		dir = filepath.Dir(dir)
	}
	if _, err := os.Stat(fileName); err == nil {
		return append(files, fileName)
	}
	return files
}

// Return the contents of the file at the given path as a string
func ReadFileAsString(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
