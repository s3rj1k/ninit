package sysinit

import (
	"context"
	"os"
	"os/exec"
	"strings"

	"github.com/s3rj1k/ninit/pkg/log/logger"
	"github.com/s3rj1k/ninit/pkg/utils"
	"golang.org/x/sys/unix"
)

func configureExecCMD(ctx context.Context, c Config, _ logger.Logger) *exec.Cmd {
	cmd := exec.CommandContext( //nolint: gosec // executing command passed from config
		ctx,
		c.GetCommandPath(),
		c.GetCommandArgs()...,
	)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if c.GetWorkDirectory() != "" {
		cmd.Dir = c.GetWorkDirectory()
	}

	cmd.Env = utils.FilterStringSlice(
		os.Environ(),
		func(x string) bool {
			return !strings.HasPrefix(x, c.GetEnvPrefix())
		},
	)

	cmd.SysProcAttr = &unix.SysProcAttr{
		// create a dedicated pidgroup for signal forwarding
		Setpgid: true,
	}

	return cmd
}

func configurePreReloadExecCMD(ctx context.Context, c Config, _ logger.Logger) *exec.Cmd {
	if c.GetPreReloadCommandPath() == "" {
		return nil
	}

	cmd := exec.CommandContext( //nolint: gosec // executing command passed from config
		ctx,
		c.GetPreReloadCommandPath(),
		c.GetPreReloadCommandArgs()...,
	)

	if c.GetWorkDirectory() != "" {
		cmd.Dir = c.GetWorkDirectory()
	}

	cmd.Env = utils.FilterStringSlice(
		os.Environ(),
		func(x string) bool {
			return !strings.HasPrefix(x, c.GetEnvPrefix())
		},
	)

	cmd.SysProcAttr = &unix.SysProcAttr{
		// create a dedicated pidgroup for signal forwarding
		Setpgid: true,
	}

	return cmd
}
