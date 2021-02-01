package ssh

import (
	"bytes"
	"fmt"
	"time"

	"golang.org/x/crypto/ssh"
)

// Client SSH Client
type Client struct {
	c *ssh.Client
}

func newSSHConfig(username string, timeout int) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		Timeout:         time.Duration(timeout) * time.Second,
		User:            username,
		Auth:            []ssh.AuthMethod{},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

// NewClient Factory
func NewClient(config *Config) (*Client, error) {
	err := config.Check()
	if err != nil {
		return nil, err
	}
	sshConfig := newSSHConfig(config.Username, config.Timeout)

	if config.Password != "" {
		sshConfig.Auth = append(sshConfig.Auth, ssh.Password(config.Password))
	}

	if config.Key != "" {
		signer, err := ssh.ParsePrivateKey([]byte(config.Key))
		if err != nil {
			return nil, err
		}
		sshConfig.Auth = append(sshConfig.Auth, ssh.PublicKeys(signer))
	}

	c, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", config.IP, config.Port), sshConfig)
	if err != nil {
		return nil, err
	}
	return &Client{c}, nil
}

// Exec execute a command
func (c *Client) Exec(command string) (stdout string, stderr string, err error) {

	session, err := c.c.NewSession()
	if err != nil {
		return stdout, stderr, err
	}
	defer session.Close()

	var stdOutBuf bytes.Buffer
	session.Stdout = &stdOutBuf

	var stdErrBuf bytes.Buffer
	session.Stderr = &stdErrBuf

	err = session.Run(command)
	if err != nil {
		return stdout, stderr, err
	}

	return string(stdOutBuf.Bytes()), string(stdErrBuf.Bytes()), err

}
