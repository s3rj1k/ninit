package hash

import (
	"io/fs"
	"path/filepath"
	"sort"
)

func getListOfFilesFromPath(path string) ([]string, error) {
	var files []string

	path = filepath.Clean(path)

	if err := filepath.WalkDir(path, func(file string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		files = append(files, filepath.ToSlash(file))

		return nil
	}); err != nil {
		return nil, err
	}

	sort.Strings(files)

	return files, nil
}
