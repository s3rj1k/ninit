package configmap

import (
	"time"
)

// Config defines package configuration interface.
type Config interface {
	GetK8sBaseDirectory() string
	GetK8sNamespace() string
	GetK8sObjectName() string
	GetPauseChannel() chan bool
	GetWatchInterval() time.Duration
}
