package sysinit

import (
	"errors"
	"time"

	"github.com/s3rj1k/ninit/pkg/log/logger"
	"golang.org/x/sys/unix"
)

func sendSignal(log logger.Logger, pid int, sig unix.Signal) {
	// forward signal to main process and all children
	if err := unix.Kill(pid, sig); err != nil {
		if errors.Is(err, unix.ESRCH) { // no such process
			log.Tracef("%v\n", err)
		} else {
			log.Warnf("%v\n", err)
		}
	}

	if sig == unix.SIGINT || sig == unix.SIGTERM {
		// lets sleep here for a bit to allow
		// application finish writing to stdout/stderr
		time.Sleep(1 * time.Second)
	}
}
