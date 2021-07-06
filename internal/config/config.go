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

var defaultConfigFile = "./config.yml"

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
func LoadConfig(configFile string, config *C) error {
	configContent, err := ioutil.ReadFile(configFile)
	if err != nil {
		errMsg := fmt.Sprintf("ioutil.ReadFile(%s) failed, error: %s", configFile, err)
		logrus.Error(errMsg)
		return errors.New(errMsg)
	}

	err = yaml.Unmarshal(configContent, &config)
	if err != nil {
		errMsg := fmt.Sprintf("yaml.Unmarshal %s failed, error: %s", configFile, err)
		logrus.Error(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

// InitConfig inits configuration
func InitConfig(configFile string) error {

	if configFile == "" {
		configFile = defaultConfigFile
	}

	Config = C{}

	err := LoadConfig(configFile, &Config)
	if err != nil {
		errMsg := fmt.Sprintf("LoadConfig(%s) failed, error: %s", configFile, err)
		logrus.Error(errMsg)
		return errors.New(errMsg)
	}

	configStr, _ := Config.ToString()
	logrus.Info(string(configStr))

	return Config.Check()

}
