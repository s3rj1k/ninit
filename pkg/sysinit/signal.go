package sysinit

import (
	"errors"
	"syscall"

	"github.com/s3rj1k/ninit/pkg/config"
)

func sendSignal(c *config.Config, pid int, sig syscall.Signal) {
	// forward signal to main process and all children
	if err := syscall.Kill(pid, sig); err != nil {
		if errors.Is(err, syscall.ESRCH) { // no such process
			c.Log.Tracef("%v\n", err)
		} else {
			c.Log.Warnf("%v\n", err)
		}
	}
}
