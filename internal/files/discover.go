package files

import (
	"os"
	"path/filepath"

	"github.com/bmatcuk/doublestar/v4"
)

func Discover(paths []string, includes []string, excludes []string) ([]string, error) {
	var matched []string

	// If no paths specified, use current directory
	if len(paths) == 0 {
		paths = []string{"."}
	}

	for _, p := range paths {
		info, err := os.Stat(p)
		if err != nil {
			return nil, err
		}

		if !info.IsDir() {
			matched = append(matched, p)
			continue
		}

		// It's a directory, walk it
		err = filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			// Check excludes first
			for _, exclude := range excludes {
				match, err := doublestar.Match(exclude, path)
				if err == nil && match {
					return nil // Skip excluded
				}
			}

			// Check includes
			for _, include := range includes {
				match, err := doublestar.Match(include, path)
				if err == nil && match {
					matched = append(matched, path)
					break
				}
			}

			return nil
		})

		if err != nil {
			return nil, err
		}
	}

	return matched, nil
}
