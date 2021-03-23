package sysinit

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/s3rj1k/ninit/pkg/config"
	"github.com/s3rj1k/ninit/pkg/reaper"
	"github.com/s3rj1k/ninit/pkg/signals"
	"github.com/s3rj1k/ninit/pkg/watcher"
)

// Run starts system init process with provided config,
// it will forward signals to child process, reap zombies,
// send reload signal on config chage.
func Run(c *config.Config, wg *sync.WaitGroup) error {
	if os.Getpid() != 1 {
		return fmt.Errorf("expecting to be run as PID 1")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	defer close(sigs)

	signal.Notify(
		sigs,
		signals.Except(syscall.SIGCHLD, syscall.SIGCLD)..., // "17", "SIGCHLD", "SIGCLD": only useful for zombie reaping
	)
	defer signal.Reset()

	cmd := configureExecCMD(ctx, c)

	if err := cmd.Start(); err != nil {
		return err //nolint: wrapcheck // error message wrapping is done by `GetErrorMessage(err error) string`
	}

	c.Log.Infof("started process '%v' with PID '%d'\n", cmd.String(), cmd.Process.Pid)

	watch := watcher.Path(ctx, wg, c.WatchPath, c.WatchInterval)
	reap := reaper.Run(ctx, wg)

	wg.Add(1)

	go func() {
		for {
			select {
			case <-ctx.Done():
				wg.Done()

				return

			case sig := <-sigs:
				signalEvent(c, sig, cmd)

			case v := <-watch:
				watcherEvent(c, v, cmd)

			case v := <-reap:
				reaperEvent(c, v)
			}
		}
	}()

	err := cmd.Wait()
	c.Log.Infof("finished process '%v' with PID '%d'\n", cmd.String(), cmd.Process.Pid)

	return err //nolint: wrapcheck // error message wrapping is done by `GetErrorMessage(err error) string`
}
