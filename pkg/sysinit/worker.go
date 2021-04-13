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

func worker(
	ctx context.Context,
	wg *sync.WaitGroup,
	c Config,
	log logger.Logger,
	cmd *exec.Cmd,
	sigs <-chan os.Signal,
	watch <-chan watcher.Message,
	reap <-chan reaper.Message,
) {
	for {
		select {
		case <-ctx.Done():
			wg.Done()

			return

		case sig := <-sigs:
			signalEvent(c, log, sig, cmd)

		case v := <-watch:
			watcherEvent(c, log, v, cmd)

		case v := <-reap:
			reaperEvent(c, log, v)
		}
	}
}
