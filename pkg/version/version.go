package version

import (
	"strings"
)

func dataOrUnknown(v string) string {
	v = strings.TrimSpace(v)

	if v == "" {
		return "UNKNOWN"
	}

	return v
}

// GetVersion returns embedded version info.
func GetVersion() string {
	return dataOrUnknown(version)
}

// GetVersion returns embedded build time.
func GetBuildTime() string {
	return dataOrUnknown(buildTime)
}
