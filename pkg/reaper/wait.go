package reaper

import (
	"errors"
	"fmt"

	"golang.org/x/sys/unix"
)

func syscallWait() (*Message, bool) {
	var status unix.WaitStatus

	// wait for orphaned zombie process
	// https://man7.org/linux/man-pages/man2/wait.2.html
	pid, err := unix.Wait4(-1, &status, unix.WNOHANG|unix.WCONTINUED, nil)

	if errors.Is(err, unix.ECHILD) {
		// no un-reaped child(ren) exist
		return &Message{
			Message: "reaper cleanup: no (more) zombies found",
		}, true
	}

	if pid == -1 {
		// error from syscall
		return &Message{
			Error: fmt.Errorf("reaper error: %w", err),
		}, false
	}

	if pid > 0 {
		// child was reaped
		return &Message{
			Message: fmt.Sprintf("reaper cleanup: pid=%d, status=%+v", pid, status),
		}, false
	}

	// `pid == 0`: one or more child(ren) exist that have not yet changed state
	return nil, false
}
