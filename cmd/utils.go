// In cmd/utils.go
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config is the main structure for the .portal.toml file
type Config struct {
	Enter PortalAction `mapstructure:"enter"`
	Leave PortalAction `mapstructure:"leave"`
}

// PortalAction defines the commands and message for an action
type PortalAction struct {
	Commands []string `mapstructure:"commands"`
	Message  string   `mapstructure:"message"`
}

// findProjectRoot searches from the current directory upwards for .portal.toml
func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, ".portal.toml")); err == nil {
			return dir, nil
		}
		if dir == "/" || dir == filepath.Dir(dir) {
			break
		}
		dir = filepath.Dir(dir)
	}
	return "", fmt.Errorf("no .portal.toml file found in this directory or any parent directory")
}

// loadConfig reads and parses the .portal.toml file from a given path
func loadConfig(path string) (*Config, error) {
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName(".portal")
	v.SetConfigType("toml")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
