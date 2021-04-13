package wlogrus

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"sync"

	"github.com/s3rj1k/ninit/pkg/capitalise"
	"github.com/s3rj1k/ninit/pkg/log/logger"
	"github.com/sirupsen/logrus"
)

// LogrusLogger is a package level logger using logrus.
type LogrusLogger struct {
	Log *logrus.Logger

	level logger.Level
	mu    sync.Mutex
}

// New creates new Logrus logger.
func New() *LogrusLogger {
	l := new(LogrusLogger)

	l.Log = logrus.New()

	l.Log.SetFormatter(
		&logrus.TextFormatter{
			DisableColors:    true,
			DisableSorting:   true,
			DisableTimestamp: true,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				pc := make([]uintptr, 20)
				n := runtime.Callers(1, pc)

				if n > 0 {
					pc = pc[:n]
					frames := runtime.CallersFrames(pc)
					next := false

					for {
						frame, more := frames.Next()

						if next {
							return fmt.Sprintf("%s()", path.Base(frame.Function)), fmt.Sprintf("%s:%d", path.Base(frame.File), frame.Line)
						}

						if f.PC == frame.PC {
							next = true
						}

						if !more {
							break
						}
					}
				}

				return "", ""
			},
		},
	)

	l.Log.SetReportCaller(true)

	return l
}

// Logf is unleveled logger.
func (l *LogrusLogger) Logf(format string, args ...interface{}) {
	if l == nil {
		panic("logger undefined")
	}

	l.Log.Logf(logrus.TraceLevel, strings.TrimSpace(capitalise.First(fmt.Sprintf(format, args...))))
}

// Tracef is a trace level logger.
func (l *LogrusLogger) Tracef(format string, args ...interface{}) {
	if l == nil {
		panic("logger undefined")
	}

	l.Log.Tracef(strings.TrimSpace(capitalise.First(fmt.Sprintf(format, args...))))
}

// Debugf is a debug level logger.
func (l *LogrusLogger) Debugf(format string, args ...interface{}) {
	if l == nil {
		panic("logger undefined")
	}

	l.Log.Debugf(strings.TrimSpace(capitalise.First(fmt.Sprintf(format, args...))))
}

// Infof is a info level logger.
func (l *LogrusLogger) Infof(format string, args ...interface{}) {
	if l == nil {
		panic("logger undefined")
	}

	l.Log.Infof(strings.TrimSpace(capitalise.First(fmt.Sprintf(format, args...))))
}

// Warnf is a warn level logger.
func (l *LogrusLogger) Warnf(format string, args ...interface{}) {
	if l == nil {
		panic("logger undefined")
	}

	l.Log.Warnf(strings.TrimSpace(capitalise.First(fmt.Sprintf(format, args...))))
}

// Errorf is a error level logger.
func (l *LogrusLogger) Errorf(format string, args ...interface{}) {
	if l == nil {
		panic("logger undefined")
	}

	l.Log.Errorf(strings.TrimSpace(capitalise.First(fmt.Sprintf(format, args...))))
}

// Fatalf is a fatal level logger.
func (l *LogrusLogger) Fatalf(format string, args ...interface{}) {
	if l == nil {
		panic("logger undefined")
	}

	l.Log.Fatalf(strings.TrimSpace(capitalise.First(fmt.Sprintf(format, args...))))
}

// Panicf is a panic level logger.
func (l *LogrusLogger) Panicf(format string, args ...interface{}) {
	if l == nil {
		panic("logger undefined")
	}

	l.Log.Panicf(strings.TrimSpace(capitalise.First(fmt.Sprintf(format, args...))))
}

// SetLevel defines maximum level of a logger output verbosity.
func (l *LogrusLogger) SetLevel(level logger.Level) {
	if l == nil {
		panic("logger undefined")
	}

	var logrusLevel logrus.Level

	switch level {
	case logger.TraceLevelLog:
		logrusLevel = logrus.TraceLevel
	case logger.DebugLevelLog:
		logrusLevel = logrus.DebugLevel
	case logger.InfoLevelLog:
		logrusLevel = logrus.InfoLevel
	case logger.WarnLevelLog:
		logrusLevel = logrus.WarnLevel
	case logger.ErrorLevelLog:
		logrusLevel = logrus.ErrorLevel
	case logger.FatalLevelLog:
		logrusLevel = logrus.FatalLevel
	case logger.PanicLevelLog:
		logrusLevel = logrus.PanicLevel
	}

	l.mu.Lock()

	l.Log.SetLevel(logrusLevel)
	l.level = level

	l.mu.Unlock()
}

// GetLevel returns current level of a logger output verbosity.
func (l *LogrusLogger) GetLevel() logger.Level {
	return l.level
}
