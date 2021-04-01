package utils

import (
	"fmt"
	"io/fs"
	"os"
)

const (
	ownerBit = 0100
	groupBit = 0010
	otherBit = 0001
	allBits  = 0111
)

// GetMode return `fs.FileMode` object or error for specified path.
func GetMode(path string) (fs.FileMode, error) {
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

// IsExecOwner returns true when filemode has exec owner bit set.
func IsExecOwner(mode os.FileMode) bool {
	return mode&ownerBit != 0
}

// IsExecGroup returns true when filemode has exec group bit set.
func IsExecGroup(mode os.FileMode) bool {
	return mode&groupBit != 0
}

// IsExecOther returns true when filemode has exec other bit set.
func IsExecOther(mode os.FileMode) bool {
	return mode&otherBit != 0
}

// IsExecAny returns true when filemode has one of exec owner, group, other bit set.
func IsExecAny(mode os.FileMode) bool {
	return mode&allBits != 0
}

// IsExecAll returns true when filemode has all of exec owner, group, other bit set.
func IsExecAll(mode os.FileMode) bool {
	return mode&allBits == allBits
}
