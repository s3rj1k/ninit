package watcher

// Message describes output from Path function.
type Message struct {
	Error     error
	Message   string
	IsChanged bool
}
