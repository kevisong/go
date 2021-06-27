package config

import (
	"errors"
	"fmt"

	"github.com/KEVISONG/go/pkg/log"
	"github.com/sirupsen/logrus"
)

// Config defines configuration
type Config struct {
	Log log.Config `json:"log" yaml:"log"`
}

// DefaultConfig is default configuration
var DefaultConfig Config

// Check checks if config is valid
func (c *Config) Check() error {
	return c.Log.Check()
}

// LoadConfig loads configuration from file
func LoadConfig(configPath string, config *Config) error {
	// TODO
	return nil
}

// InitConfig inits configuration
func InitConfig(configPath string) error {

	err := LoadConfig(configPath, &DefaultConfig)
	if err != nil {
		errMsg := fmt.Sprintf("LoadConfig(%s) failed, error: %s", configPath, err)
		logrus.Error(errMsg)
		return errors.New(errMsg)
	}

	return DefaultConfig.Check()

}
