package watcher

import (
	"fmt"
	"time"
)

// Message describes output from Path function.
type Message struct {
	Error     error
	Message   string
	IsChanged bool
}

func paused(path string) Message {
	return Message{
		Message: fmt.Sprintf("path '%s' watch paused", path),
	}
}

func resumed(path string) Message {
	return Message{
		Message: fmt.Sprintf("path '%s' watch resumed", path),
	}
}

func change(path string, delta time.Duration) Message {
	return Message{
		IsChanged: true,
		Message:   fmt.Sprintf("path '%s' change detected (%v)", path, delta),
	}
}

func hashError(path string, err error) Message {
	return Message{
		Error: fmt.Errorf("hash compute error, path '%s': %w", path, err),
	}
}

func shutdown(path string) Message {
	return Message{
		Message: fmt.Sprintf("path '%s' watch is shutting down", path),
	}
}
