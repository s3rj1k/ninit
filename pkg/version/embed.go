package version

import _ "embed" // go embed blank import

//nolint:gochecknoglobals // go embed fields
var (
	//go:embed .autogenerated/version
	version string
	//go:embed .autogenerated/buildTime
	buildTime string
)