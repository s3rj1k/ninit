package utils

import (
	"os"
	"path/filepath"
)

// IsHelpFlag is a KISS CLI `--help` or `-h` arguments checker.
func IsHelpFlag() bool {
	if len(os.Args) > 1 && // os.Args[0] contains application name
		(os.Args[1] == "-h" || os.Args[1] == "--help") {
		return true
	}

	return false
}

// GetApplicationName just a conviniece wrapper
// to get running application base name.
func GetApplicationName() string {
	return filepath.Base(os.Args[0])
}
