package sysinit

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/s3rj1k/ninit/pkg/config"
	"github.com/s3rj1k/ninit/pkg/reaper"
	"github.com/s3rj1k/ninit/pkg/watcher"
)

func signalEvent(c *config.Config, sig os.Signal, cmd *exec.Cmd) {
	if sig == nil || sig == syscall.SIGCHLD {
		return
	}

	if v, ok := sig.(syscall.Signal); ok {
		pid := -cmd.Process.Pid
		if c.SignalToDirectChildOnly {
			pid = cmd.Process.Pid
		}

		sendSignal(c, pid, v)

		c.Log.Debugf("sent '%v' signal to PID '%d'\n", sig, -cmd.Process.Pid) // can be very verbose
	}
}

func watcherEvent(c *config.Config, v watcher.Message, cmd *exec.Cmd) {
	if v.Error != nil {
		c.Log.Errorf("%v\n", v.Error)
	}

	if v.Message != "" {
		c.Log.Infof("%v\n", v.Message)
	}

	if v.IsChanged {
		pid := cmd.Process.Pid
		if c.ReloadSignalToPGID {
			pid = -cmd.Process.Pid
		}

		sendSignal(c, pid, c.ReloadSignal)

		c.Log.Infof("sent '%v' signal to PID '%d'\n", c.ReloadSignal, pid)
	}
}

func reaperEvent(c *config.Config, v reaper.Message) {
	if v.Error != nil {
		c.Log.Errorf("%v\n", v.Error)
	}

	if v.Message != "" {
		c.Log.Infof("%v\n", v.Message)
	}
}
