package standart

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/s3rj1k/ninit/pkg/capitalise"
	"github.com/s3rj1k/ninit/pkg/log/logger"
)

// Standart is a package level logger.
type Standart struct {
	LogLevel *log.Logger

	TraceLevel *log.Logger
	DebugLevel *log.Logger
	InfoLevel  *log.Logger
	WarnLevel  *log.Logger
	ErrorLevel *log.Logger
	FatalLevel *log.Logger

	output io.Writer
	prefix string

	level logger.Level
	mu    sync.Mutex
}

const (
	// DefaultFlags defines default flags for standart logger.
	DefaultFlags = log.Lmsgprefix

	callDepth = 2
)

// Create creates new logger.
func Create(out io.Writer, prefix string, flags int, level logger.Level) *Standart {
	l := &Standart{
		output: out,
		prefix: prefix,
		level:  level,

		LogLevel: log.New(
			out,
			prefix+"[LOG]: ",
			flags),
		TraceLevel: log.New(
			ioutil.Discard,
			prefix+"[TRACE]: ",
			flags),
		DebugLevel: log.New(
			ioutil.Discard,
			prefix+"[DEBUG]: ",
			flags),
		InfoLevel: log.New(
			ioutil.Discard,
			prefix+"[INFO]: ",
			flags),
		WarnLevel: log.New(
			ioutil.Discard,
			prefix+"[WARN]: ",
			flags),
		ErrorLevel: log.New(
			ioutil.Discard,
			prefix+"[ERROR]: ",
			flags),
		FatalLevel: log.New(
			ioutil.Discard,
			prefix+"[FATAL]: ",
			flags),
	}

	l.SetLevel(level)

	return l
}

// Logf is unleveled logger.
func (l *Standart) Logf(format string, args ...interface{}) {
	if l == nil {
		panic("logger undefined")
	}

	if l.LogLevel == nil {
		panic("logger undefined")
	}

	if err := l.LogLevel.Output(callDepth, capitalise.First(fmt.Sprintf(format, args...))); err != nil {
		panic(err)
	}
}

// Tracef is a trace level logger.
func (l *Standart) Tracef(format string, args ...interface{}) {
	if l == nil {
		panic("logger undefined")
	}

	if l.TraceLevel == nil {
		panic("trace logger undefined")
	}

	if err := l.TraceLevel.Output(callDepth, capitalise.First(fmt.Sprintf(format, args...))); err != nil {
		panic(err)
	}
}

// Debugf is a debug level logger.
func (l *Standart) Debugf(format string, args ...interface{}) {
	if l == nil {
		panic("logger undefined")
	}

	if l.DebugLevel == nil {
		panic("debug logger undefined")
	}

	if err := l.DebugLevel.Output(callDepth, capitalise.First(fmt.Sprintf(format, args...))); err != nil {
		panic(err)
	}
}

// Infof is a info level logger.
func (l *Standart) Infof(format string, args ...interface{}) {
	if l == nil {
		panic("logger undefined")
	}

	if l.InfoLevel == nil {
		panic("info logger undefined")
	}

	if err := l.InfoLevel.Output(callDepth, capitalise.First(fmt.Sprintf(format, args...))); err != nil {
		panic(err)
	}
}

// Warnf is a warn level logger.
func (l *Standart) Warnf(format string, args ...interface{}) {
	if l == nil {
		panic("logger undefined")
	}

	if l.InfoLevel == nil {
		panic("warn logger undefined")
	}

	if err := l.WarnLevel.Output(callDepth, capitalise.First(fmt.Sprintf(format, args...))); err != nil {
		panic(err)
	}
}

// Errorf is a error level logger.
func (l *Standart) Errorf(format string, args ...interface{}) {
	if l == nil {
		panic("logger undefined")
	}

	if l.ErrorLevel == nil {
		panic("error logger undefined")
	}

	if err := l.ErrorLevel.Output(callDepth, capitalise.First(fmt.Sprintf(format, args...))); err != nil {
		panic(err)
	}
}

//revive:disable:deep-exit // Fatalf LogLevel needs to call `os.Exit()`.

// Fatalf is a fatal level logger.
func (l *Standart) Fatalf(format string, args ...interface{}) {
	if l == nil {
		panic("logger undefined")
	}

	if l.FatalLevel == nil {
		panic("fatal logger undefined")
	}

	if err := l.FatalLevel.Output(callDepth, capitalise.First(fmt.Sprintf(format, args...))); err != nil {
		panic(err)
	}

	os.Exit(1)
}

//revive:enable:deep-exit

// Panicf is a panic level logger.
func (l *Standart) Panicf(format string, args ...interface{}) {
	s := l.prefix + ": " + capitalise.First(fmt.Sprintf(format, args...))

	panic(s)
}

func (l *Standart) checkDefinedLoggers() error {
	if l == nil {
		return fmt.Errorf("logger undefined")
	}

	if l.TraceLevel == nil {
		return fmt.Errorf("trace logger undefined")
	}

	if l.DebugLevel == nil {
		return fmt.Errorf("debug logger undefined")
	}

	if l.InfoLevel == nil {
		return fmt.Errorf("info logger undefined")
	}

	if l.WarnLevel == nil {
		return fmt.Errorf("warn logger undefined")
	}

	if l.ErrorLevel == nil {
		return fmt.Errorf("error logger undefined")
	}

	if l.FatalLevel == nil {
		return fmt.Errorf("fatal logger undefined")
	}

	return nil
}

// SetLevel defines maximum level of a logger output verbosity.
func (l *Standart) SetLevel(level logger.Level) {
	if err := l.checkDefinedLoggers(); err != nil {
		panic(err)
	}

	l.mu.Lock()

	l.level = level

	if level >= logger.TraceLevelLog {
		l.TraceLevel.SetOutput(l.output)
	}

	if level >= logger.DebugLevelLog {
		l.DebugLevel.SetOutput(l.output)
	}

	if level >= logger.InfoLevelLog {
		l.InfoLevel.SetOutput(l.output)
	}

	if level >= logger.WarnLevelLog {
		l.WarnLevel.SetOutput(l.output)
	}

	if level >= logger.ErrorLevelLog {
		l.ErrorLevel.SetOutput(l.output)
	}

	if level >= logger.FatalLevelLog {
		l.FatalLevel.SetOutput(l.output)
	}

	l.mu.Unlock()
}

// GetLevel returns current level of a logger output verbosity.
func (l *Standart) GetLevel() logger.Level {
	return l.level
}
