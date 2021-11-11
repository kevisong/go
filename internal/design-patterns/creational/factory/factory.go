package factory

import (
	"fmt"
)

type iGun interface {
	SetName(string)
	GetName() string
	SetPower(int)
	GetPower() int
}

type Gun struct {
	Name  string
	Power int
}

func (g *Gun) SetName(name string) {
	g.Name = name
}

func (g *Gun) GetName() string {
	return g.Name
}

func (g *Gun) SetPower(power int) {
	g.Power = power
}

func (g *Gun) GetPower() int {
	return g.Power
}

type AK struct {
	Gun
}

func NewAK() iGun {
	return &AK{}
}

type M4 struct {
	Gun
}

func NewM4() iGun {
	return &M4{}
}

// NewGun is the factory method
func NewGun(gunType string) (iGun, error) {
	switch gunType {
	case "AK":
		return NewAK(), nil
	case "M4":
		return NewM4(), nil
	default:
		return nil, fmt.Errorf("gunType %s not supported", gunType)
	}
}
