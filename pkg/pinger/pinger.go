package pinger

import (
	"time"

	"github.com/go-ping/ping"
)

// Pinger Pinger
type Pinger struct {
	pinger *ping.Pinger
}

// NewPinger NewPinger
func NewPinger() (*Pinger, error) {
	pinger, err := ping.NewPinger("127.0.0.1")
	if err != nil {
		return nil, err
	}
	return &Pinger{
		pinger: pinger,
	}, nil
}

// SetAddr SetAddr
func (p *Pinger) SetAddr(addr string) *Pinger {
	if p.pinger != nil {
		p.pinger.SetAddr(addr)
	}
	return p
}

// SetCount SetCount
func (p *Pinger) SetCount(n int) *Pinger {
	if p.pinger != nil {
		p.pinger.Count = n
	}
	return p
}

// SetTimeout SetTimeout
func (p *Pinger) SetTimeout(n int) *Pinger {
	if p.pinger != nil {
		p.pinger.Timeout = time.Duration(n)
	}
	return p
}

// Run Run
func (p *Pinger) Run() error {
	return p.pinger.Run()
}

// Statistics Statistics
func (p *Pinger) Statistics() *ping.Statistics {
	return p.pinger.Statistics()
}
