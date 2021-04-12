package sysinit

import (
	"time"

	"golang.org/x/sys/unix"
)

// Config defines package configuration interface.
type Config interface {
	GetCommandArgs() []string
	GetCommandPath() string
	GetEnvPrefix() string
	GetPauseChannel() chan bool
	GetReloadSignal() unix.Signal
	GetReloadSignalToPGID() bool
	GetSignalToDirectChildOnly() bool
	GetWatchInterval() time.Duration
	GetWatchPath() string
	GetWorkDirectory() string
}
