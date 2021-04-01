package shared

// Specifies default package level constants.
const (
	DefaultEnvPrefix = "INIT_"
	DefaultLogPrefix = "init "

	DefaultWatchIntervalInSeconds = 3
	NanosecondsInSeconds          = 1000 * 1000 * 1000

	UnknownValue = "UNKNOWN"
)

// DescriptionPrefix defines message prefix for CLI help.
const DescriptionPrefix = `
Application:
	Name: %NAME%
	Version: %VERSION%
	Build Time: %BUILD_TIME%
`
