package klog

import (
	"fmt"
	"sync"

	log "github.com/s3rj1k/ninit/pkg/log/logger"
)

//nolint: gochecknoglobals // Is needed for package level logger.
var (
	logger log.Logger
	mu     = sync.Mutex{}
)

// https://github.com/kubernetes/community/blob/master/contributors/devel/sig-instrumentation/logging.md
const (
	disableLevelThreshold = 0
	infoLevelThreshold    = 3
	debugLevelThreshold   = 5
)

// SetLogger redirects klog logging to the given logger.
// It must be called prior any call to klogger.
func SetLogger(l log.Logger) {
	mu.Lock()
	logger = l
	mu.Unlock()
}

type Level int32

func V(level Level) Level {
	return level
}

func (l Level) Enabled() bool {
	level := logger.GetLevel()

	switch {
	case l < disableLevelThreshold:
		return false
	case l < infoLevelThreshold:
		return level >= log.InfoLevelLog
	case l < debugLevelThreshold:
		return level >= log.DebugLevelLog
	default:
		return level >= log.TraceLevelLog
	}
}

func (l Level) Info(args ...interface{}) {
	switch {
	case l < disableLevelThreshold:
		return
	case l < infoLevelThreshold:
		logger.Infof(fmt.Sprint(args...))
	case l < debugLevelThreshold:
		logger.Debugf(fmt.Sprint(args...))
	default:
		logger.Tracef(fmt.Sprint(args...))
	}
}

func (l Level) Infoln(args ...interface{}) {
	switch {
	case l < disableLevelThreshold:
		return
	case l < infoLevelThreshold:
		logger.Infof(fmt.Sprint(args...), "\n")
	case l < debugLevelThreshold:
		logger.Debugf(fmt.Sprint(args...), "\n")
	default:
		logger.Tracef(fmt.Sprint(args...), "\n")
	}
}

func (l Level) Infof(format string, args ...interface{}) {
	switch {
	case l < disableLevelThreshold:
		return
	case l < infoLevelThreshold:
		logger.Infof(format, args...)
	case l < debugLevelThreshold:
		logger.Debugf(format, args...)
	default:
		logger.Tracef(format, args...)
	}
}

func Info(args ...interface{}) {
	logger.Logf(fmt.Sprint(args...))
}

func InfoDepth(_ int, args ...interface{}) {
	logger.Logf(fmt.Sprint(args...))
}

func Infoln(args ...interface{}) {
	logger.Logf(fmt.Sprint(args...), "\n")
}

func Infof(format string, args ...interface{}) {
	logger.Logf(format, args...)
}

func Warning(args ...interface{}) {
	logger.Warnf(fmt.Sprint(args...))
}

func WarningDepth(_ int, args ...interface{}) {
	logger.Warnf(fmt.Sprint(args...))
}

func Warningln(args ...interface{}) {
	logger.Warnf(fmt.Sprint(args...), "\n")
}

func Warningf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Error(args ...interface{}) {
	logger.Errorf(fmt.Sprint(args...))
}

func ErrorDepth(_ int, args ...interface{}) {
	logger.Errorf(fmt.Sprint(args...))
}

func Errorln(args ...interface{}) {
	logger.Errorf(fmt.Sprint(args...), "\n")
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatalf(fmt.Sprint(args...))
}

func FatalDepth(_ int, args ...interface{}) {
	logger.Fatalf(fmt.Sprint(args...))
}

func Fatalln(args ...interface{}) {
	logger.Fatalf(fmt.Sprint(args...), "\n")
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

func Exit(args ...interface{}) {
	logger.Fatalf(fmt.Sprint(args...))
}

func ExitDepth(_ int, args ...interface{}) {
	logger.Fatalf(fmt.Sprint(args...))
}

func Exitln(args ...interface{}) {
	logger.Fatalf(fmt.Sprint(args...), "\n")
}

func Exitf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}
