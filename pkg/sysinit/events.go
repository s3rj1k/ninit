package sysinit

import (
	"os"
	"os/exec"

	"github.com/s3rj1k/ninit/pkg/log/logger"
	"github.com/s3rj1k/ninit/pkg/reaper"
	"github.com/s3rj1k/ninit/pkg/watcher"
	"golang.org/x/sys/unix"
)

func signalEvent(c Config, log logger.Logger, sig os.Signal, cmd *exec.Cmd) {
	if sig == nil {
		return
	}

	signal, ok := sig.(unix.Signal)
	if !ok || signal == unix.SIGCHLD {
		return
	}

	pid := -cmd.Process.Pid
	if c.GetSignalToDirectChildOnly() {
		pid = cmd.Process.Pid
	}

	sendSignal(log, pid, signal)

	log.Debugf("sent '%v' signal to PID '%d'\n", sig, -cmd.Process.Pid) // can be very verbose
}

func watcherEvent(c Config, log logger.Logger, v watcher.Message, cmd, preReloadCmd *exec.Cmd) {
	if v.Error != nil {
		log.Errorf("%v\n", v.Error)
	}

	if v.Message != "" {
		log.Infof("%v\n", v.Message)
	}

	if v.IsChanged {
		pid := cmd.Process.Pid
		if c.GetReloadSignalToPGID() {
			pid = -cmd.Process.Pid
		}

		if preReloadCmd != nil {
			log.Debugf("pre-reload command defined: %s\n", preReloadCmd.String())

			if err := preReloadCmd.Run(); err != nil {
				log.Errorf("failed to send '%v' signal, pre-reload command failed: %v\n", c.GetReloadSignal(), err)

				return
			}
		}

		sendSignal(log, pid, c.GetReloadSignal())

		log.Infof("sent '%v' signal to PID '%d'\n", c.GetReloadSignal(), pid)
	}
}

func reaperEvent(_ Config, log logger.Logger, v reaper.Message) {
	if v.Error != nil {
		log.Errorf("%v\n", v.Error)
	}

	if v.Message != "" {
		log.Infof("%v\n", v.Message)
	}
}
