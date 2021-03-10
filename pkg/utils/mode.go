package utils

import (
	"os"
)

// IsExecOwner returns true when filemode has exec owner bit set.
func IsExecOwner(mode os.FileMode) bool {
	return mode&0100 != 0
}

// IsExecGroup returns true when filemode has exec group bit set.
func IsExecGroup(mode os.FileMode) bool {
	return mode&0010 != 0
}

// IsExecOther returns true when filemode has exec other bit set.
func IsExecOther(mode os.FileMode) bool {
	return mode&0001 != 0
}

// IsExecAny returns true when filemode has one of exec owner, group, other bit set.
func IsExecAny(mode os.FileMode) bool {
	return mode&0111 != 0
}

// IsExecAll returns true when filemode has all of exec owner, group, other bit set.
func IsExecAll(mode os.FileMode) bool {
	return mode&0111 == 0111
}
