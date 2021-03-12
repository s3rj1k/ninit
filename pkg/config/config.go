package config

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/s3rj1k/ninit/pkg/signals"
	"github.com/s3rj1k/ninit/pkg/utils"
)

// Specifies default prefixes for environment variables and logs.
const (
	DefaultEnvPrefix = "INIT_"
	DefaultLogPrefix = "init: "
)

// Config contains application configuration.
type Config struct {
	CommandPath   string
	WatchPath     string
	WorkDirectory string

	ReloadSignal  syscall.Signal
	WatchInterval time.Duration

	CommandArgs []string

	// contains application specific prefix for environment variables
	EnvPrefix string
}

// Help prints user-friendly help with available configuration options.
func Help(prefix, name, version, buildTime string) {
	text := `
	Application:
		Name: %NAME%
		Version: %VERSION%
		Build Time: %BUILD_TIME%

	Avaliable envars configuration options:
		- %PREFIX%COMMAND_PATH
				path to executable [required]
		- %PREFIX%COMMAND_ARGS
				command arguments
		- %PREFIX%WORK_DIRECTORY_PATH
				path to application new current working directory
		- %PREFIX%RELOAD_SIGNAL
				OS signal what triggers application config reload [default 'SIGHUP']
		- %PREFIX%WATCH_INTERVAL
				watch (type: pulling) time interval [default '3s']
		- %PREFIX%WATCH_PATH
				file or directory path to watch (type: pulling) file changes recursevely
	`
	if name == "" {
		name = filepath.Base(os.Args[0])
	}

	if version == "" {
		version = "UNKNOWN"
	}

	if buildTime == "" {
		buildTime = "UNKNOWN"
	}

	r := strings.NewReplacer(
		"\t", "  ",
		"%PREFIX%", prefix,
		"%NAME%", name,
		"%VERSION%", version,
		"%BUILD_TIME%", buildTime,
	)

	fmt.Println(r.Replace(strings.TrimPrefix(text, "\n")))
}

// setCommandPath reads command path from `prefix + COMMAND_PATH` env.
func (c *Config) setCommandPath(prefix string) error {
	env := prefix + "COMMAND_PATH"

	c.CommandPath = os.Getenv(env)

	if strings.TrimSpace(c.CommandPath) == "" {
		return fmt.Errorf("path is invalid, empty string")
	}

	mode, err := getMode(c.CommandPath)
	if err != nil {
		return fmt.Errorf("%s: %w", env, err)
	}

	if mode.IsDir() {
		return fmt.Errorf("%s: path '%s' is directory", env, c.CommandPath)
	}

	if !mode.IsRegular() {
		return fmt.Errorf("%s: path '%s' is not regular file", env, c.CommandPath)
	}

	if !utils.IsExecOwner(mode) {
		return fmt.Errorf("%s: path '%s' has no exec owner bit set", env, c.CommandPath)
	}

	if !utils.IsExecGroup(mode) {
		return fmt.Errorf("%s: path '%s' has no exec group bit set", env, c.CommandPath)
	}

	return nil
}

// setWatchPath reads watch path from `prefix + WATCH_PATH` env.
func (c *Config) setWatchPath(prefix string) error {
	env := prefix + "WATCH_PATH"

	c.WatchPath = os.Getenv(env)

	if strings.TrimSpace(c.WatchPath) == "" {
		c.WatchPath = ""

		return nil
	}

	mode, err := getMode(c.WatchPath)
	if err != nil {
		return fmt.Errorf("%s: %w", env, err)
	}

	if !mode.IsDir() && !mode.IsRegular() {
		return fmt.Errorf("%s: path '%s' is not file or directory", env, c.WatchPath)
	}

	return nil
}

// setWorkingDirectory reads working directory path from `prefix + WORK_DIRECTORY_PATH` env.
func (c *Config) setWorkingDirectory(prefix string) error {
	env := prefix + "WORK_DIRECTORY_PATH"

	c.WorkDirectory = os.Getenv(env)

	if strings.TrimSpace(c.WorkDirectory) == "" {
		c.WorkDirectory = ""

		return nil
	}

	mode, err := getMode(c.WorkDirectory)
	if err != nil {
		return fmt.Errorf("%s: %w", env, err)
	}

	if !mode.IsDir() {
		return fmt.Errorf("%s: path '%s' is not directory", env, c.WorkDirectory)
	}

	return nil
}

// setCommandArgs reads command args from `prefix + COMMAND_ARGS` env.
func (c *Config) setCommandArgs(prefix string) error {
	env := prefix + "COMMAND_ARGS"
	c.CommandArgs = strings.Fields(os.Getenv(env))

	return nil
}

// setWatchInterval reads pulling interval from `prefix + WATCH_INTERVAL` env.
func (c *Config) setWatchInterval(prefix string) error {
	env := prefix + "WATCH_INTERVAL"
	val := os.Getenv(env)
	if strings.TrimSpace(val) == "" {
		// c.WatchInterval, _ = time.ParseDuration("3s")
		c.WatchInterval = 3 * 1000 * 1000 * 1000 // 3 seconds in nanoseconds

		return nil
	}

	t, err := time.ParseDuration(val)
	if err != nil || t < 0 {
		return fmt.Errorf("%s: invalid time duration '%s'", env, val)
	}

	c.WatchInterval = t

	return nil
}

// setReloadSignal reads reload signal from `prefix + RELOAD_SIGNAL` env.
func (c *Config) setReloadSignal(prefix string) error {
	env := prefix + "RELOAD_SIGNAL"
	val := os.Getenv(env)
	if strings.TrimSpace(val) == "" {
		c.ReloadSignal = syscall.SIGHUP

		return nil
	}

	sig, err := signals.Parse(val)
	if err != nil {
		return fmt.Errorf("%s: %v", env, err)
	}

	c.ReloadSignal = sig

	return nil
}

// validateEnvPrefix runs basic sanity check on environment variable prefix.
func validateEnvPrefix(prefix string) error {
	re := "^[A-Z_]{1,}[A-Z0-9_]{0,}_$"
	if !regexp.MustCompile(re).MatchString(prefix) {
		return fmt.Errorf("envars prefix '%s' must match '%s' regexp", prefix, re)
	}

	return nil
}

// Get returns validated configuration object filled with default values or from environment variables.
func Get(prefix string) (*Config, error) {
	c := new(Config)

	if err := validateEnvPrefix(prefix); err != nil {
		return nil, err
	}

	c.EnvPrefix = prefix

	if err := c.setCommandPath(prefix); err != nil {
		return nil, err
	}
	if err := c.setCommandArgs(prefix); err != nil {
		return nil, err
	}
	if err := c.setWorkingDirectory(prefix); err != nil {
		return nil, err
	}
	if err := c.setWatchPath(prefix); err != nil {
		return nil, err
	}
	if err := c.setWatchInterval(prefix); err != nil {
		return nil, err
	}
	if err := c.setReloadSignal(prefix); err != nil {
		return nil, err
	}

	return c, nil
}
