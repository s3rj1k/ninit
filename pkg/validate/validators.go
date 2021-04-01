package validate

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/s3rj1k/ninit/pkg/signals"
	"github.com/s3rj1k/ninit/pkg/utils"
	"golang.org/x/sys/unix"
)

// EnvironName validates environment variable name.
func EnvironName(name string) error {
	re := "^[A-Z_]{1,}[A-Z0-9_]{0,}$"
	if !regexp.MustCompile(re).MatchString(name) {
		return fmt.Errorf("environ name '%s' must match '%s' regexp", name, re)
	}

	return nil
}

// Executable validate that path is valid executable.
func Executable(path string) error {
	if strings.TrimSpace(path) == "" {
		return fmt.Errorf("path is invalid, empty string")
	}

	mode, err := utils.GetMode(path)
	if err != nil {
		return err //nolint: wrapcheck // error string formed in external package is styled correctly
	}

	if mode.IsDir() {
		return fmt.Errorf("path '%s' is directory", path)
	}

	if !mode.IsRegular() {
		return fmt.Errorf("path '%s' is not regular file", path)
	}

	if !utils.IsExecOwner(mode) {
		return fmt.Errorf("path '%s' has no exec owner bit set", path)
	}

	if !utils.IsExecGroup(mode) {
		return fmt.Errorf("path '%s' has no exec group bit set", path)
	}

	if err := unix.Access(filepath.Clean(path), unix.R_OK); err != nil {
		return fmt.Errorf("path '%s' is not readable", path)
	}

	if err := unix.Access(filepath.Clean(path), unix.X_OK); err != nil {
		return fmt.Errorf("path '%s' is not executable", path)
	}

	return nil
}

// FileOrDirectory validate that path is valid (readable/writable) regular file or directory.
func FileOrDirectory(path string) error {
	if strings.TrimSpace(path) == "" {
		return fmt.Errorf("path is invalid, empty string")
	}

	mode, err := utils.GetMode(path)
	if err != nil {
		return err //nolint: wrapcheck // error string formed in external package is styled correctly
	}

	filepath.Clean(path)

	if !mode.IsDir() && !mode.IsRegular() {
		return fmt.Errorf("path '%s' is not file or directory", path)
	}

	if err := unix.Access(filepath.Clean(path), unix.R_OK); err != nil {
		return fmt.Errorf("path '%s' is not readable", path)
	}

	if err := unix.Access(filepath.Clean(path), unix.W_OK); err != nil {
		return fmt.Errorf("path '%s' is not writable", path)
	}

	return nil
}

// Directory validate that path is valid directory.
func Directory(path string) error {
	if strings.TrimSpace(path) == "" {
		return fmt.Errorf("path is invalid, empty string")
	}

	mode, err := utils.GetMode(path)
	if err != nil {
		return err //nolint: wrapcheck // error string formed in external package is styled correctly
	}

	if !mode.IsDir() {
		return fmt.Errorf("path '%s' is not directory", path)
	}

	if err := unix.Access(filepath.Clean(path), unix.R_OK); err != nil {
		return fmt.Errorf("path '%s' is not readable", path)
	}

	if err := unix.Access(filepath.Clean(path), unix.W_OK); err != nil {
		return fmt.Errorf("path '%s' is not writable", path)
	}

	return nil
}

// Duration validate that value is parsable `time.Duration`.
func Duration(val string) error {
	t, err := time.ParseDuration(val)
	if err != nil || t < 0 {
		return fmt.Errorf("invalid time duration '%s'", val)
	}

	return nil
}

// Bool validate that value is valid Bool (true/false).
func Bool(val string) error {
	if strings.EqualFold(val, "true") || strings.EqualFold(val, "false") {
		return nil
	}

	return fmt.Errorf("invalid bool value '%s', can be only 'true' or 'false'", val)
}

// Signal validate that value is valid signal name.
func Signal(val string) error {
	_, err := signals.Parse(val)
	if err != nil {
		return err //nolint: wrapcheck // error string formed in external package is styled correctly
	}

	return nil
}

// DNSLabel validate that value is valid DNS label based on RFC 1123.
//  * https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#dns-label-names
//  * https://tools.ietf.org/html/rfc1123
func DNSLabel(value string) error {
	re := `^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?$`
	if !regexp.MustCompile(re).MatchString(value) {
		return fmt.Errorf("value '%s' must match '%s' regexp (DNS label, RFC 1123)", value, re)
	}

	return nil
}
