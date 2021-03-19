package sysinit

import (
	"fmt"
	"os/exec"
)

func GetErrorMessage(err error) string {
	if err == nil {
		return ""
	}

	if exitErr, ok := err.(*exec.ExitError); ok {
		return fmt.Sprintf("command exited with code: %d, %v", exitErr.ExitCode(), err)
	}

	return err.Error()
}
