package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

const DefaultConfigFile = "./config/default.toml"

// LoadConfig loads the configuration from the default TOML file into the given struct.
func LoadConfig(config interface{}) error {
	v := viper.New()
	v.SetConfigFile(DefaultConfigFile)
	v.SetConfigType("toml")
	v.AutomaticEnv() // Allow environment variables to override config values

	// Read the config file
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("error loading config file: %w", err)
	}

	// Unmarshal into the provided struct
	if err := v.Unmarshal(config); err != nil {
		return fmt.Errorf("error parsing config file: %w", err)
	}

	log.Printf("Config loaded from %s", DefaultConfigFile)
	return nil
}
