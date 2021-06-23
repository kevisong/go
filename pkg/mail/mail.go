package mail

import (
	gomail "gopkg.in/gomail.v2"
)

// Sender defines the email sender
type Sender struct {
	d *gomail.Dialer
}

// NewSender factory
func NewSender(host string, port int, username, password string) *Sender {
	return &Sender{
		d: gomail.NewDialer(host, port, username, password),
	}
}

// Send sends the email
func (s *Sender) Send(from, to, subject, body string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	if err := s.d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
