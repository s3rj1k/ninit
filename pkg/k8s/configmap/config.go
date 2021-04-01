package configmap

import (
	"time"
)

// Config defines package configuration interface.
type Config interface {
	GetK8sBaseDirectory() string
	GetK8sObjectName() string
	GetK8sNamespace() string
	GetWatchInterval() time.Duration
}
