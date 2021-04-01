package logger

type Level int32

// Available log levels.
const (
	PanicLevelLog Level = iota + 1
	FatalLevelLog
	ErrorLevelLog
	WarnLevelLog
	InfoLevelLog
	DebugLevelLog
	TraceLevelLog
)

// Logger defines custom logger interface.
type Logger interface {
	// unleveled logger
	Logf(format string, args ...interface{})

	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	SetLevel(level Level)
	GetLevel() Level
}
