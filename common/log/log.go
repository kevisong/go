package log

import (
	"fmt"
	"os"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"

	"github.com/sirupsen/logrus"
)

// Logrus initialization pkg
// github.com/sirupsen/logrus

// Config for log config
type Config struct {
	Level           string `json:"level" yaml:"level"`                       // Log level
	Path            string `json:"path" yaml:"path"`                         // Log path
	Filename        string `json:"filename" yaml:"filename"`                 // Log filename
	MaxAge          int    `json:"max_age" yaml:"max_age"`                   // Log store time
	RotationTime    int    `json:"rotation_time" yaml:"rotation_time"`       // Log rotation time
	Formatter       string `json:"formatter" yaml:"formatter"`               // Log formatter type
	ReportCaller    bool   `json:"report_caller" yaml:"report_caller"`       // Report caller func and line number
	TimestampFormat string `json:"timestamp_format" yaml:"timestamp_format"` // Timestamp Format
	LogrusLevel     logrus.Level
	rMaxAge         time.Duration
	rRotationTime   time.Duration
}

func (c *Config) checkLevel() error {
	c.Level = strings.TrimSpace(c.Level)
	if c.Level == "" {
		c.Level = "info"
	}
	logrusLevel, err := logrus.ParseLevel(c.Level)
	if err != nil {
		return err
	}
	c.LogrusLevel = logrusLevel
	return nil
}
func (c *Config) checkPath() {
	c.Path = strings.TrimSpace(c.Path)
	if c.Path == "" {
		c.Path = "./"
	}
}

func (c *Config) checkMaxAge() error {
	if c.MaxAge <= 0 {
		c.MaxAge = 30 * 24 * 3600
	}
	rMaxAge, err := time.ParseDuration(fmt.Sprintf("%ds", c.MaxAge))
	if err != nil {
		return err
	}
	c.rMaxAge = rMaxAge
	return nil
}

func (c *Config) checkRotationTime() error {
	if c.RotationTime <= 0 {
		c.RotationTime = 24 * 3600
	}
	rRotationTime, err := time.ParseDuration(fmt.Sprintf("%ds", c.RotationTime))
	if err != nil {
		return err
	}
	c.rRotationTime = rRotationTime
	return nil
}

func (c *Config) checkFormatter() error {
	c.Formatter = strings.ToUpper(strings.TrimSpace(c.Formatter))
	if c.Formatter == "" {
		c.Formatter = "TEXT"
	}
	if c.Formatter != "TEXT" && c.Formatter != "JSON" {
		return fmt.Errorf("not a valid log formatter: %s", c.Formatter)
	}
	return nil
}

func (c *Config) checkTimestampFormat() {
	c.TimestampFormat = strings.TrimSpace(c.TimestampFormat)
	if c.TimestampFormat == "" {
		c.TimestampFormat = "2006-01-02 15:04:05.000000"
	}
}

// Check validates config
func (c *Config) Check() error {

	// Check level
	err := c.checkLevel()
	if err != nil {
		return err
	}

	// Check output path
	c.checkPath()

	c.Filename = strings.TrimSpace(c.Filename)

	// Check max age
	err = c.checkMaxAge()
	if err != nil {
		return err
	}

	// Check rotation time
	err = c.checkRotationTime()
	if err != nil {
		return err
	}

	// Check formatter
	err = c.checkFormatter()
	if err != nil {
		return err
	}

	// Check timestamp format
	c.checkTimestampFormat()

	return nil

}

func setLogrus(c *Config) {
	switch c.Formatter {
	case "TEXT":
		formatter := &logrus.TextFormatter{}
		formatter.TimestampFormat = c.TimestampFormat
		logrus.SetFormatter(formatter)

	case "JSON":
		formatter := &logrus.JSONFormatter{}
		formatter.TimestampFormat = c.TimestampFormat
		logrus.SetFormatter(formatter)
	}
	logrus.SetReportCaller(c.ReportCaller)
	logrus.SetLevel(c.LogrusLevel)
}

func setRotatelogs(c *Config) error {
	logRotate, err := rotatelogs.New(
		fmt.Sprintf("%s.%%Y%%m%%d%%H%%M", fmt.Sprintf("%s/%s", c.Path, c.Filename)),
		rotatelogs.WithLinkName(c.Filename),
		rotatelogs.WithMaxAge(c.rMaxAge),
		rotatelogs.WithRotationTime(c.rRotationTime),
	)
	if err != nil {
		errorInfo := fmt.Sprintf("failed to create rotatelogs: %s", err)
		fmt.Println(errorInfo)
		return err
	}

	logrus.SetOutput(logRotate)
	return nil
}

// Init initialize logrus to your weapon of choice
func Init(c *Config) error {

	err := c.Check()
	if err != nil {
		return err
	}

	setLogrus(c)

	// Output to cli
	if c.Filename == "" {
		logrus.SetOutput(os.Stdout)
		return nil
	}

	// Output to file
	return setRotatelogs(c)
}
