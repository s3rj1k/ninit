package logger

// Available log levels.
const (
	PanicLevelLog int = iota + 1
	FatalLevelLog
	ErrorLevelLog
	WarnLevelLog
	InfoLevelLog
	DebugLevelLog
	TraceLevelLog
)

// Logger defines custom logger interface.
type Logger interface {
	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	SetLevel(level int)
}
