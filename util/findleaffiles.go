package util

import (
	"strings"
	"os"
	"path/filepath"
	"github.com/rs/zerolog/log"
)

func FindLeafFiles(search string, includes []string, excludes []string) ([]string, error) {
	var files []string
	var leafs []string
	var err error

	for _, dir := range includes {
		err = filepath.Walk(dir, func(fullpath string, stat os.FileInfo, err error) error {
			if err != nil {
				log.Error().Msgf("%v: %v", fullpath, err)
				return err
			}
			if stat.Name() == search {
				files = append(files, fullpath)
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	outer:
		for _, check := range files {
			dir := filepath.Dir(check)
			if (dir == "." || dir == "/") && len(files) > 1 {
				continue
			}

			// deduplicate - skip already detected leafs
			for _, leaf := range leafs {
				if leaf == check {
					continue outer
				}
			}

			for _, file := range files {
				if file == check {
					continue
				}
				if strings.HasPrefix(file, dir) {
					continue outer
				}
			}

			leafs = append(leafs, check)
		}

	return leafs, nil
}
