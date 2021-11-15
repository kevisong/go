package builder

import "fmt"

type Car struct {
	Seats  int
	Engine int
	GPS    int
}

// NewCar has a problem: the constructor parameters can be very long
func NewCar(seats, engine, gps int) *Car {
	return &Car{Seats: seats, Engine: engine, GPS: gps}
}

type Builder interface {
	Reset()
	SetSeats(int)
	SetEngine(int)
	SetGPS(int)
	GetResult() Car
}

// SportsCarBuilder is the concrete builder to build the car step by step
type SportsCarBuilder struct {
	// car is the result
	car *Car
}

func (s SportsCarBuilder) Reset() {
	s.car = &Car{}
}
func (s SportsCarBuilder) SetSeats(i int) {
	// Install seats
}
func (c SportsCarBuilder) SetEngine(i int) {
	// Install engine
}
func (s SportsCarBuilder) SetGPS(i int) {
	// Install GPS
}
func (s SportsCarBuilder) GetResult() Car {
	return *s.car
}

type SportsCarBuilderDirector struct {
	builder Builder
}

func (s *SportsCarBuilderDirector) SetBuilder(builder Builder) {
	s.builder = builder
}

func (s *SportsCarBuilderDirector) Build() Car {
	s.builder.SetSeats(2)
	s.builder.SetEngine(10)
	s.builder.SetGPS(0)
	return s.builder.GetResult()
}

func Run() {
	builder := SportsCarBuilder{}
	director := SportsCarBuilderDirector{}
	director.SetBuilder(builder)
	sportsCar := director.Build()
	fmt.Println(sportsCar)
}
