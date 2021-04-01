package sysinit

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"sync"

	"github.com/s3rj1k/ninit/pkg/logger"
	"github.com/s3rj1k/ninit/pkg/reaper"
	"github.com/s3rj1k/ninit/pkg/signals"
	"github.com/s3rj1k/ninit/pkg/watcher"
	"golang.org/x/sys/unix"
)

// Run starts and waits for termination of system init process
// with provided config, it will forward signals to child process,
// reap zombies, send reload signal on config chage.
func Run(ctx context.Context, wg *sync.WaitGroup, c Config, log logger.Logger) error {
	if os.Getpid() != 1 {
		return fmt.Errorf("expecting to be run as PID 1")
	}

	sigs := make(chan os.Signal, 1)
	defer close(sigs)

	signal.Notify(
		sigs,
		signals.Except(unix.SIGCHLD, unix.SIGCLD)..., // "17", "SIGCHLD", "SIGCLD": only useful for zombie reaping
	)
	defer signal.Reset()

	cmd := configureExecCMD(ctx, c, log)

	if err := cmd.Start(); err != nil {
		return err //nolint: wrapcheck // error message wrapping is done by `GetErrorMessage(err error) string`
	}

	log.Infof("started process '%v' with PID '%d'\n", cmd.String(), cmd.Process.Pid)

	watch := watcher.Path(ctx, wg, c.GetWatchPath(), c.GetWatchInterval())
	reap := reaper.Run(ctx, wg)

	wg.Add(1)

	go func(
		ctx context.Context,
		wg *sync.WaitGroup,
		c Config,
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
	}(ctx, wg, c, cmd, sigs, watch, reap)

	err := cmd.Wait()
	log.Infof("finished process '%v' with PID '%d'\n", cmd.String(), cmd.Process.Pid)

	return err //nolint: wrapcheck // error message wrapping is done by `GetErrorMessage(err error) string`
}
