package logger

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Standart is a package level logger.
type Standart struct {
	TraceLevel *log.Logger
	DebugLevel *log.Logger
	InfoLevel  *log.Logger
	WarnLevel  *log.Logger
	ErrorLevel *log.Logger
	FatalLevel *log.Logger
	PanicLevel *log.Logger
}

// DefaultFlags defines default flags for standart logger.
const DefaultFlags = log.Lmsgprefix

// Create creates new logger.
func Create(prefix string, flags, level int) *Standart {
	l := &Standart{
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
		PanicLevel: log.New(
			ioutil.Discard,
			prefix+"PANIC: ",
			flags),
	}
	l.SetLevel(level)

	return l
}

// Trace is a trace level logger
func (l *Standart) Tracef(format string, args ...interface{}) {
	if l == nil {
		panic("logger undefined")
	}

	if l.TraceLevel == nil {
		panic("trace logger undefined")
	}

	if err := l.TraceLevel.Output(2, fmt.Sprintf(format, args...)); err != nil {
		panic(err)
	}
}

// Debug is a debug level logger
func (l *Standart) Debugf(format string, args ...interface{}) {
	if l == nil {
		panic("logger undefined")
	}

	if l.DebugLevel == nil {
		panic("debug logger undefined")
	}

	if err := l.DebugLevel.Output(2, fmt.Sprintf(format, args...)); err != nil {
		panic(err)
	}
}

// Info is a info level logger
func (l *Standart) Infof(format string, args ...interface{}) {
	if l == nil {
		panic("logger undefined")
	}

	if l.InfoLevel == nil {
		panic("info logger undefined")
	}

	if err := l.InfoLevel.Output(2, fmt.Sprintf(format, args...)); err != nil {
		panic(err)
	}
}

// Warn is a warn level logger
func (l *Standart) Warnf(format string, args ...interface{}) {
	if l == nil {
		panic("logger undefined")
	}

	if l.InfoLevel == nil {
		panic("warn logger undefined")
	}

	if err := l.WarnLevel.Output(2, fmt.Sprintf(format, args...)); err != nil {
		panic(err)
	}
}

// Error is a error level logger
func (l *Standart) Errorf(format string, args ...interface{}) {
	if l == nil {
		panic("logger undefined")
	}

	if l.ErrorLevel == nil {
		panic("error logger undefined")
	}

	if err := l.ErrorLevel.Output(2, fmt.Sprintf(format, args...)); err != nil {
		panic(err)
	}
}

//revive:disable:deep-exit

// Fatal is a fatal level logger
func (l *Standart) Fatalf(format string, args ...interface{}) {
	if l == nil {
		panic("logger undefined")
	}

	if l.FatalLevel == nil {
		panic("fatal logger undefined")
	}

	if err := l.FatalLevel.Output(2, fmt.Sprintf(format, args...)); err != nil {
		panic(err)
	}

	os.Exit(1)
}

//revive:disable:deep-exit

// Panic is a panic level logger
func (l *Standart) Panicf(format string, args ...interface{}) {
	if l == nil {
		panic("logger undefined")
	}

	if l.PanicLevel == nil {
		panic("panic logger undefined")
	}

	s := fmt.Sprintf(format, args...)

	if err := l.PanicLevel.Output(2, s); err != nil {
		panic(err)
	}

	panic(s)
}

// SetLevel defines maximum level of a logger output
func (l *Standart) SetLevel(level int) {
	if l == nil {
		panic("logger undefined")
	}

	if l.TraceLevel == nil {
		panic("trace logger undefined")
	}

	if l.DebugLevel == nil {
		panic("debug logger undefined")
	}

	if l.InfoLevel == nil {
		panic("info logger undefined")
	}

	if l.WarnLevel == nil {
		panic("warn logger undefined")
	}

	if l.ErrorLevel == nil {
		panic("error logger undefined")
	}

	if l.FatalLevel == nil {
		panic("fatal logger undefined")
	}

	if l.PanicLevel == nil {
		panic("panic logger undefined")
	}

	if level >= TraceLevelLog {
		l.TraceLevel.SetOutput(os.Stderr)
	}

	if level >= DebugLevelLog {
		l.DebugLevel.SetOutput(os.Stderr)
	}

	if level >= InfoLevelLog {
		l.InfoLevel.SetOutput(os.Stderr)
	}

	if level >= WarnLevelLog {
		l.WarnLevel.SetOutput(os.Stderr)
	}

	if level >= ErrorLevelLog {
		l.ErrorLevel.SetOutput(os.Stderr)
	}

	if level >= FatalLevelLog {
		l.FatalLevel.SetOutput(os.Stderr)
	}

	if level >= PanicLevelLog {
		l.PanicLevel.SetOutput(os.Stderr)
	}
}
