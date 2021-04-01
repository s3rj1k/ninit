package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/s3rj1k/ninit/pkg/config/shared"
	"github.com/s3rj1k/ninit/pkg/signals"
	"github.com/s3rj1k/ninit/pkg/validate"
	"golang.org/x/sys/unix"
)

const DescriptionBody = `
Available envars configuration options:
	- %PREFIX%COMMAND_PATH
			path to executable [required].
	- %PREFIX%COMMAND_ARGS
			command arguments.
	- %PREFIX%WORK_DIRECTORY_PATH
			path to application new current working directory.

	- %PREFIX%RELOAD_SIGNAL
			OS signal what triggers application config reload [default 'SIGHUP'].
	- %PREFIX%RELOAD_SIGNAL_TO_PGID
			send reload signal to PGID instead of PID.
	- %PREFIX%SIGNAL_TO_DIRECT_CHILD_ONLY
			signals (excluding reload signal) are only forwarded
			to direct child and not to any of its descendants,
			meaning signal is sent to PID instead of PGID.

	- %PREFIX%WATCH_INTERVAL
			watch (type: pulling) time interval [default '3s'].
	- %PREFIX%WATCH_PATH
			file or directory path to watch (type: pulling) file changes recursevely.
`

// Redefine defaults from shared package for convenient importing.
const (
	DefaultEnvPrefix = shared.DefaultEnvPrefix
	DefaultLogPrefix = shared.DefaultLogPrefix
)

// Config contains application configuration.
type Config struct {
	envPrefix string // contains application specific prefix for environment variables

	watchPath string

	commandPath   string
	workDirectory string
	commandArgs   []string

	reloadSignal  unix.Signal
	watchInterval time.Duration

	signalToDirectChildOnly bool
	reloadSignalToPGID      bool
}

// New creates new config with defaul values.
func New(prefix string) *Config {
	return &Config{
		envPrefix:     prefix,
		reloadSignal:  unix.SIGHUP,
		watchInterval: shared.DefaultWatchIntervalInSeconds * shared.NanosecondsInSeconds,
	}
}

func (c *Config) Help(name, version, buildTime string) {
	shared.Help(name, version, buildTime, c.GetEnvPrefix(), c.GetDescriptionBody())
}

func (*Config) GetDefaultEnvPrefix() string { return shared.DefaultEnvPrefix }
func (*Config) GetDefaultLogPrefix() string { return shared.DefaultLogPrefix }
func (*Config) GetDescriptionBody() string  { return DescriptionBody }

func (c *Config) GetCommandArgs() []string         { return c.commandArgs }
func (c *Config) GetCommandPath() string           { return c.commandPath }
func (c *Config) GetEnvPrefix() string             { return c.envPrefix }
func (c *Config) GetReloadSignal() unix.Signal     { return c.reloadSignal }
func (c *Config) GetReloadSignalToPGID() bool      { return c.reloadSignalToPGID }
func (c *Config) GetSignalToDirectChildOnly() bool { return c.signalToDirectChildOnly }
func (c *Config) GetWatchInterval() time.Duration  { return c.watchInterval }
func (c *Config) GetWatchPath() string             { return c.watchPath }
func (c *Config) GetWorkDirectory() string         { return c.workDirectory }

// Get reads environment variables to update and validate configuration object.
func (c *Config) Get() error {
	if err := c.SetCommandPath("COMMAND_PATH"); err != nil {
		return err
	}

	if err := c.SetCommandArgs("COMMAND_ARGS"); err != nil {
		return err
	}

	if err := c.SetWorkingDirectory("WORK_DIRECTORY_PATH"); err != nil {
		return err
	}

	if err := c.SetWatchPath("WATCH_PATH"); err != nil {
		return err
	}

	if err := c.SetWatchInterval("WATCH_INTERVAL"); err != nil {
		return err
	}

	if err := c.SetReloadSignal("RELOAD_SIGNAL"); err != nil {
		return err
	}

	if err := c.SetReloadSignalToPGID("RELOAD_SIGNAL_TO_PGID"); err != nil {
		return err
	}

	return c.SetSignalToDirectChildOnly("SIGNAL_TO_DIRECT_CHILD_ONLY")
}

// SetCommandPath reads command path from environ and updates its value inside config.
func (c *Config) SetCommandPath(env string) error {
	env = c.envPrefix + env

	val, _, err := shared.LookupEnvValue(env)
	if err != nil {
		return err //nolint: wrapcheck // error string formed in external package is styled correctly
	}

	if err := validate.Executable(val); err != nil {
		return fmt.Errorf("%s: %w", env, err)
	}

	c.commandPath = val

	return nil
}

// SetWatchPath reads watch path from environ and updates its value inside config.
func (c *Config) SetWatchPath(env string) error {
	env = c.envPrefix + env

	val, ok, err := shared.LookupEnvValue(env)
	if err != nil {
		return err //nolint: wrapcheck // error string formed in external package is styled correctly
	}

	if !ok {
		return nil
	}

	err = validate.FileOrDirectory(val)
	if err != nil {
		return fmt.Errorf("%s: %w", env, err)
	}

	c.watchPath = val

	return nil
}

// SetWorkingDirectory reads working directory path from environ and updates its value inside config.
func (c *Config) SetWorkingDirectory(env string) error {
	env = c.envPrefix + env

	val, ok, err := shared.LookupEnvValue(env)
	if err != nil {
		return err //nolint: wrapcheck // error string formed in external package is styled correctly
	}

	if !ok {
		return nil
	}

	err = validate.Directory(val)
	if err != nil {
		return fmt.Errorf("%s: %w", env, err)
	}

	c.workDirectory = val

	return nil
}

// SetCommandArgs reads command args from environ and updates its value inside config.
func (c *Config) SetCommandArgs(env string) error {
	env = c.envPrefix + env

	val, _, err := shared.LookupEnvValue(env)
	if err != nil {
		return err //nolint: wrapcheck // error string formed in external package is styled correctly
	}

	c.commandArgs = strings.Fields(val)

	return nil
}

// SetWatchInterval reads pulling interval from environ and updates its value inside config.
func (c *Config) SetWatchInterval(env string) error {
	env = c.envPrefix + env

	val, ok, err := shared.LookupEnvValue(env)
	if err != nil {
		return err //nolint: wrapcheck // error string formed in external package is styled correctly
	}

	if !ok {
		return nil
	}

	err = validate.Duration(val)
	if err != nil {
		return fmt.Errorf("%s: %w", env, err)
	}

	c.watchInterval, _ = time.ParseDuration(val)

	return nil
}

// SetReloadSignalToPGID reads bool value from environ and updates its value inside config.
func (c *Config) SetReloadSignalToPGID(env string) error {
	env = c.envPrefix + env

	val, ok, err := shared.LookupEnvValue(env)
	if err != nil {
		return err //nolint: wrapcheck // error string formed in external package is styled correctly
	}

	if !ok {
		return nil
	}

	err = validate.Bool(val)
	if err != nil {
		return fmt.Errorf("%s: %w", env, err)
	}

	if strings.EqualFold(val, "true") {
		c.reloadSignalToPGID = true
	}

	return nil
}

// SetSignalToDirectChildOnly reads bool value from environ and updates its value inside config.
func (c *Config) SetSignalToDirectChildOnly(env string) error {
	env = c.envPrefix + env

	val, ok, err := shared.LookupEnvValue(env)
	if err != nil {
		return err //nolint: wrapcheck // error string formed in external package is styled correctly
	}

	if !ok {
		return nil
	}

	err = validate.Bool(val)
	if err != nil {
		return fmt.Errorf("%s: %w", env, err)
	}

	if strings.EqualFold(val, "true") {
		c.signalToDirectChildOnly = true
	}

	return nil
}

// SetReloadSignal reads reload signal from environ and updates its value inside config.
func (c *Config) SetReloadSignal(env string) error {
	env = c.envPrefix + env

	val, ok, err := shared.LookupEnvValue(env)
	if err != nil {
		return err //nolint: wrapcheck // error string formed in external package is styled correctly
	}

	if !ok {
		return nil
	}

	err = validate.Signal(val)
	if err != nil {
		return fmt.Errorf("%s: %w", env, err)
	}

	c.reloadSignal, _ = signals.Parse(val)

	return nil
}
