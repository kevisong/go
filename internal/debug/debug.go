package debug

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Config Config
type Config struct {
	Port int `yaml:"port"`
}

// Run runs debug server
func Run(c Config) {
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", c.Port), nil)
	if err != nil {
		logrus.Error(err)
		return
	}
}
