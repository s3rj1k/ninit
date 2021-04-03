package config

import (
	"fmt"
	"strings"

	cfg "github.com/s3rj1k/ninit/pkg/config/minimal"
	"github.com/s3rj1k/ninit/pkg/config/shared"
	"github.com/s3rj1k/ninit/pkg/validate"
)

const DescriptionBody = `
	- %PREFIX%K8S_BASE_DIRECTORY_PATH
			base directory path to apply kubernetes ConfigMaps based on received event:
				- ADDED, MODIFIED: file content is written to directory %PREFIX%K8S_BASE_DIRECTORY_PATH,
					files are named based on KEY values from ConfigMap Data and BinaryData sections.
				- DELETED: all regular files inside %PREFIX%K8S_BASE_DIRECTORY_PATH are deleted.
	- %PREFIX%K8S_NAMESPACE
			specifies kubernetes namespace that contains object to watch.
	- %PREFIX%K8S_CONFIG_MAP_NAME
			specifies kubernetes object (ConfigMap) name.
`

// Redefine defaults from shared package for convenient importing.
const (
	DefaultEnvPrefix = cfg.DefaultEnvPrefix
	DefaultLogPrefix = cfg.DefaultLogPrefix
)

// Config contains application configuration.
type Config struct {
	k8sBaseDirectory string
	k8sObjectName    string
	k8sNamespace     string

	cfg.Config
}

// New creates new config with defaul values.
func New(prefix string) *Config {
	return &Config{
		Config: *cfg.New(prefix),
	}
}

func (c *Config) Help(name, version, buildTime string) {
	shared.Help(name, version, buildTime, c.GetEnvPrefix(), c.GetDescriptionBody())
}

func (*Config) GetDescriptionBody() string {
	return strings.TrimPrefix(cfg.DescriptionBody, "\n") + "\n" + strings.TrimPrefix(DescriptionBody, "\n")
}

func (c *Config) GetK8sBaseDirectory() string { return c.k8sBaseDirectory }
func (c *Config) GetK8sNamespace() string     { return c.k8sNamespace }
func (c *Config) GetK8sObjectName() string    { return c.k8sObjectName }

// Get reads environment variables to update and validate configuration object.
func (c *Config) Get() error {
	if err := c.Config.Get(); err != nil {
		return err //nolint: wrapcheck // error string formed in external package is styled correctly
	}

	if err := c.SetK8sBaseDirectory("K8S_BASE_DIRECTORY_PATH"); err != nil {
		return err
	}

	if err := c.SetK8sNamespace("K8S_NAMESPACE"); err != nil {
		return err
	}

	return c.SetK8sObjectName("K8S_CONFIG_MAP_NAME")
}

// SetK8sBaseDirectory reads k8s base directory path from environ and updates its value inside config.
func (c *Config) SetK8sBaseDirectory(env string) error {
	env = c.GetEnvPrefix() + env

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

	c.k8sBaseDirectory = val

	return nil
}

// SetK8sNamespace reads k8s namespace value from environ and updates its value inside config.
func (c *Config) SetK8sNamespace(env string) error {
	env = c.GetEnvPrefix() + env

	val, _, err := shared.LookupEnvValue(env)
	if err != nil {
		return err //nolint: wrapcheck // error string formed in external package is styled correctly
	}

	err = validate.DNSLabel(val)
	if err != nil {
		return fmt.Errorf("%s: %w", env, err)
	}

	c.k8sNamespace = val

	return nil
}

// SetK8sObjectName reads k8s object name value from environ and updates its value inside config.
func (c *Config) SetK8sObjectName(env string) error {
	env = c.GetEnvPrefix() + env

	val, _, err := shared.LookupEnvValue(env)
	if err != nil {
		return err //nolint: wrapcheck // error string formed in external package is styled correctly
	}

	err = validate.DNSLabel(val)
	if err != nil {
		return fmt.Errorf("%s: %w", env, err)
	}

	c.k8sObjectName = val

	return nil
}
