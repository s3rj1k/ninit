package sysinit

import (
	"context"
	"os"
	"os/exec"
	"sync"

	"github.com/s3rj1k/ninit/pkg/log/logger"
	"github.com/s3rj1k/ninit/pkg/reaper"
	"github.com/s3rj1k/ninit/pkg/watcher"
)

type workerConfig struct {
	cmd          *exec.Cmd
	preReloadCmd *exec.Cmd

	sigs  <-chan os.Signal
	watch <-chan watcher.Message
	reap  <-chan reaper.Message
}

func worker(
	ctx context.Context,
	wg *sync.WaitGroup,
	c Config,
	log logger.Logger,
	wc *workerConfig,
) {
	for {
		select {
		case <-ctx.Done():
			wg.Done()

			return

		case sig := <-wc.sigs:
			signalEvent(c, log, sig, wc.cmd)

		case v := <-wc.watch:
			watcherEvent(c, log, v, wc.cmd, wc.preReloadCmd)

		case v := <-wc.reap:
			reaperEvent(c, log, v)
		}
	}
}
