package log

import (
	"fmt"
	"os"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/mattn/go-colorable"
	"github.com/tietang/go-utils"
	lpf "github.com/x-cray/logrus-prefixed-formatter"

	"github.com/sirupsen/logrus"
)

// log initialization based on github.com/sirupsen/logrus

// Config for log config
type Config struct {
	Level        string `json:"level" yaml:"level"`                 // Log level
	Path         string `json:"path" yaml:"path"`                   // Log path
	Filename     string `json:"filename" yaml:"filename"`           // Log filename
	MaxAge       int    `json:"max_age" yaml:"max_age"`             // Log store time
	RotationTime int    `json:"rotation_time" yaml:"rotation_time"` // Log rotation time
	ReportCaller bool   `json:"report_caller" yaml:"report_caller"` // Report caller func and line number
	LogrusLevel  logrus.Level
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
	return nil
}

func (c *Config) checkRotationTime() error {
	if c.RotationTime <= 0 {
		c.RotationTime = 24 * 3600
	}
	return nil
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

	return nil

}

func newFormatter() *lpf.TextFormatter {
	formatter := &lpf.TextFormatter{}
	formatter.ForceColors = false
	formatter.DisableColors = true
	formatter.ForceFormatting = true
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02.15:04:05.000000"
	return formatter
}

func setHooks() {
	hook := utils.NewLineNumLogrusHook()
	hook.EnableFileNameLog = true
	hook.EnableFuncNameLog = true
	logrus.AddHook(hook)
}

func setLogrus(c *Config) {

	logrus.SetFormatter(newFormatter())
	logrus.SetOutput(colorable.NewColorableStdout())
	logrus.SetReportCaller(c.ReportCaller)
	logrus.SetLevel(c.LogrusLevel)

	setHooks()
}

func setRotatelogs(c *Config) error {

	maxAge, err := time.ParseDuration(fmt.Sprintf("%ds", c.MaxAge))
	if err != nil {
		return err
	}

	rotationTime, err := time.ParseDuration(fmt.Sprintf("%ds", c.RotationTime))
	if err != nil {
		return err
	}

	logRotate, err := rotatelogs.New(
		fmt.Sprintf("%s.%%Y%%m%%d%%H%%M", fmt.Sprintf("%s/%s", c.Path, c.Filename)),
		rotatelogs.WithLinkName(c.Filename),
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		errMsg := fmt.Sprintf("failed to create rotatelogs: %s", err)
		fmt.Println(errMsg)
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
