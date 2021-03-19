package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/s3rj1k/ninit/pkg/config"
	"github.com/s3rj1k/ninit/pkg/logger"
	"github.com/s3rj1k/ninit/pkg/reaper"
	"github.com/s3rj1k/ninit/pkg/signals"
	"github.com/s3rj1k/ninit/pkg/utils"
	"github.com/s3rj1k/ninit/pkg/version"
	"github.com/s3rj1k/ninit/pkg/watch"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		fmt.Println(
			config.Help(
				config.DefaultEnvPrefix,
				version.GetApplicationName(),
				version.GetVersion(),
				version.GetBuildTime(),
			),
		)
		os.Exit(0)
	}

	c := config.New(
		logger.Create(
			config.DefaultLogPrefix,
			logger.DefaultFlags, // for debug purposes can be set to: 'log.Lmsgprefix | log.Lshortfile | log.Lmsgprefix'
			logger.TraceLevelLog,
		),
	)

	if os.Getpid() != 1 {
		c.Log.Fatalf("expecting to be run as PID 1, exiting\n")
	}

	if err := c.Get(config.DefaultEnvPrefix); err != nil {
		c.Log.Fatalf("%v\n", err)
	}

	var wg sync.WaitGroup

	if err := run(c, &wg); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			c.Log.Errorf("command exited with code: %d, %v\n", exitErr.ExitCode(), err)
		} else {
			c.Log.Errorf("unexpected error: %v\n", err)
		}
	}

	wg.Wait()
}

func run(c *config.Config, wg *sync.WaitGroup) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	defer close(sigs)

	signal.Notify(
		sigs,
		signals.Except(syscall.SIGCHLD, syscall.SIGCLD)..., // "17", "SIGCHLD", "SIGCLD": only useful for zombie reaping
	)
	defer signal.Reset()

	cmd := exec.CommandContext(
		ctx,
		c.CommandPath,
		c.CommandArgs...,
	)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if c.WorkDirectory != "" {
		cmd.Dir = c.WorkDirectory
	}

	cmd.Env = utils.FilterStringSlice(
		os.Environ(),
		func(x string) bool {
			return !strings.HasPrefix(x, c.EnvPrefix)
		},
	)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		// create a dedicated pidgroup for signal forwarding
		Setpgid: true,
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	c.Log.Infof("started process '%v' with PID '%d'\n", cmd.String(), cmd.Process.Pid)

	watch := watch.Path(ctx, wg, c.WatchPath, c.WatchInterval)
	reap := reaper.Run(ctx, wg)

	wg.Add(1)

	go func() {
		sendSignal := func(pid int, sig syscall.Signal) {
			// forward signal to main process and all children
			if err := syscall.Kill(pid, sig); err != nil {
				if err != syscall.ESRCH { // no such process
					c.Log.Warnf("%v\n", err)
				}

				return
			}
		}

		for {
			select {
			case <-ctx.Done():
				wg.Done()
				return

			case sig := <-sigs:
				if sig == nil || sig == syscall.SIGCHLD {
					continue
				}

				if v, ok := sig.(syscall.Signal); ok {
					sendSignal(-cmd.Process.Pid, v)

					c.Log.Debugf("sent '%v' signal to PID '%d'\n", sig, -cmd.Process.Pid) // can be very verbose
				}

			case v := <-watch:
				if v.Error != nil {
					c.Log.Errorf("%v\n", v.Error)
				}

				if v.Message != "" {
					c.Log.Infof("%v\n", v.Message)
				}

				if v.IsChanged {
					var pid int = cmd.Process.Pid

					if c.ReloadSignalToPGID {
						pid = -cmd.Process.Pid
					}

					sendSignal(pid, c.ReloadSignal)

					c.Log.Infof("sent '%v' signal to PID '%d'\n", c.ReloadSignal, pid)
				}

			case v := <-reap:
				if v.Error != nil {
					c.Log.Errorf("%v\n", v.Error)
				}

				if v.Message != "" {
					c.Log.Infof("%v\n", v.Message)
				}
			}
		}
	}()

	return cmd.Wait()
}
