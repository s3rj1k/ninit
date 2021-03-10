package config

import (
	"fmt"
	"io/fs"
	"os"
)

func getMode(path string) (fs.FileMode, error) {
	info, err := os.Lstat(path)
	if os.IsPermission(err) {
		return 0, fmt.Errorf("path '%s' access denied", path)
	}

	if os.IsNotExist(err) {
		return 0, fmt.Errorf("path '%s' does not exist", path)
	}

	if info == nil {
		return 0, fmt.Errorf("path '%s' is not valid", path)
	}

	return info.Mode(), nil
}
