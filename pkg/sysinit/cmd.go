package sysinit

import (
	"context"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/s3rj1k/ninit/pkg/config"
	"github.com/s3rj1k/ninit/pkg/utils"
)

func configureExecCMD(ctx context.Context, c *config.Config) *exec.Cmd {
	cmd := exec.CommandContext( //nolint: gosec // executing command passed from config
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

	return cmd
}
