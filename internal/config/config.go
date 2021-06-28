package config

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/KEVISONG/go/pkg/log"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// C defines configuration
type C struct {
	Log log.Config `json:"log" yaml:"log"`
}

var defaultconfigPath = "./config.yml"

// DefaultConfig is default configuration
var Config C

// Check checks if config is valid
func (c *C) Check() error {
	return c.Log.Check()
}

// ToString converts server config to string
func (c *C) ToString() (string, error) {
	configBytes, err := yaml.Marshal(c)
	if err != nil {
		errMsg := fmt.Sprintf("yaml.Marshal(%v) failed, error: %s", *c, err)
		logrus.Error(errMsg)
		return "", errors.New(errMsg)
	}
	return string(configBytes), nil
}

// LoadConfig loads configuration from file
func LoadConfig(configPath string, config *C) error {
	configContent, err := ioutil.ReadFile(configPath)
	if err != nil {
		errMsg := fmt.Sprintf("ioutil.ReadFile(%s) failed, error: %s", configPath, err)
		logrus.Error(errMsg)
		return errors.New(errMsg)
	}

	err = yaml.Unmarshal(configContent, &config)
	if err != nil {
		errMsg := fmt.Sprintf("yaml.Unmarshal %s failed, error: %s", configPath, err)
		logrus.Error(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

// InitConfig inits configuration
func InitConfig(configPath string) error {

	if configPath == "" {
		configPath = defaultconfigPath
	}

	Config = C{}

	err := LoadConfig(configPath, &Config)
	if err != nil {
		errMsg := fmt.Sprintf("LoadConfig(%s) failed, error: %s", configPath, err)
		logrus.Error(errMsg)
		return errors.New(errMsg)
	}

	configStr, _ := Config.ToString()
	logrus.Info(string(configStr))

	return Config.Check()

}
