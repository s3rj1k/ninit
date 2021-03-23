package sysinit

import (
	"errors"
	"fmt"
	"os/exec"
)

func GetErrorMessage(err error) string {
	if err == nil {
		return ""
	}

	var exitErr *exec.ExitError

	if errors.As(err, &exitErr) {
		return fmt.Sprintf("command exited with code: %d, %v", exitErr.ExitCode(), err)
	}

	return err.Error()
}
