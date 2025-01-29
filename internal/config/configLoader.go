package config

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
)

// Default options for configuration loading.
const (
	DefaultConfigType     = "toml"
	DefaultConfigDir      = "./config"
	DefaultConfigFileName = "default"
	KeyWorkDirEnv         = "WORKDIR"
	AppModeTest           = "test"
)

// Options struct holds config options.
type Options struct {
	configType            string
	configPath            string
	defaultConfigFileName string
}

// Loader is a wrapper over the config loader implementation.
type Loader struct {
	opts  Options
	viper *viper.Viper
}

// NewDefaultOptions returns default options for config loading.
func NewDefaultOptions() Options {
	var configDir string

	workDir := os.Getenv(KeyWorkDirEnv)
	if workDir != "" {
		// used in containers: expects $WORKDIR/config
		configDir = path.Join(workDir, DefaultConfigDir)
	} else {
		// used in development or local environments
		configDir = DefaultConfigDir
	}

	return NewOptions(DefaultConfigType, configDir, DefaultConfigFileName)
}

// NewOptions creates new Options struct.
func NewOptions(
	configType string,
	configPath string,
	defaultConfigFileName string) Options {

	return Options{configType, configPath, defaultConfigFileName}
}

// NewDefaultLoader returns a new config loader with default options.
func NewDefaultLoader() *Loader {
	return NewLoader(NewDefaultOptions())
}

// NewLoader returns a new Loader instance.
func NewLoader(opts Options) *Loader {
	return &Loader{opts, viper.New()}
}

// Load reads environment-specific configurations and defaults, and unmarshals them into the config interface.
func (c *Loader) Load(env string, mode string, config interface{}) error {
	fmt.Println("Environment Variable:", os.Getenv("WEATHER_API_KEY"))

	// Load the default file then override it with the env-specific file
	err := c.loadByConfigName(c.opts.defaultConfigFileName, config)
	if err != nil {
		return err
	}
	return c.loadByConfigName(fmt.Sprintf("%s_%s", env, mode), config)
}

// loadByConfigName loads a configuration file and unmarshals it into the config interface.
func (c *Loader) loadByConfigName(configName string, config interface{}) error {
	if configName == DefaultConfigFileName {
		fmt.Printf(
			"Loading default config file: %v/%v.%v\n",
			c.opts.configPath,
			configName,
			c.opts.configType,
		)
	} else {
		fmt.Printf(
			"Loading config file: %v/%v.%v\n",
			c.opts.configPath,
			configName,
			c.opts.configType,
		)
	}

	c.viper.SetConfigName(configName)
	c.viper.SetConfigType(c.opts.configType)
	c.viper.AddConfigPath(c.opts.configPath)
	c.viper.AutomaticEnv()
	c.viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := c.viper.ReadInConfig(); err != nil {
		return err
	}

	return c.viper.Unmarshal(config)
}
