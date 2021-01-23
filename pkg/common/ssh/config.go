package ssh

import "errors"

// Config SSH Config
type Config struct {
	IP       string `json:"ip"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Key      string `json:"key,omitempty"`
	Timeout  int    `json:"timeout"`
}

func (c *Config) checkIP() error {
	if c.IP == "" {
		return errors.New("empty ip")
	}
	return nil
}

func (c *Config) checkPort() error {
	if c.Port <= 0 {
		return errors.New("invalid port")
	}
	return nil
}

func (c *Config) checkAuth() error {
	if c.Password == "" || c.Key == "" {
		return errors.New("empty authentication")
	}
	return nil
}

// Check config
func (c *Config) Check() error {

	err := c.checkIP()
	if err != nil {
		return err
	}

	err = c.checkPort()
	if err != nil {
		return err
	}

	err = c.checkAuth()
	if err != nil {
		return err
	}

	return nil

}
