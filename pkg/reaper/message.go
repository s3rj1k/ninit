package reaper

// Message describes output from Run function.
type Message struct {
	Error   error
	Message string
}
