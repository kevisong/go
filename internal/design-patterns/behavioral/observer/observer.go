package observer

import "fmt"

// The observer pattern is a software design pattern in which an object,
// named the subject, maintains a list of its dependents, called observers,
// and notifies them automatically of any state changes, usually by
// calling one of their methods.

type Subject interface {
	Register(Observer)
	Unregister(Observer)
	Notify(string)
}

type Observer interface {
	Name() string
	Update(string)
}

// Central is the subject that notifies all observers when
// there is a state change, which in this case - new msg.
type Central struct {
	monitors map[string]Observer
}

func (c *Central) Register(o Observer) {
	c.monitors[o.Name()] = o
}

func (c *Central) Unregister(o Observer) {
	delete(c.monitors, o.Name())
}

func (c *Central) Notify(msg string) {
	for _, o := range c.monitors {
		o.Update(msg)
	}
}

// Monitor1 is an observer
type Monitor1 struct{}

func (m *Monitor1) Name() string {
	return "monitor1"
}

func (m *Monitor1) Update(msg string) {
	fmt.Printf("monitor 1 updates %s\n", msg)
}

// Monitor2 is an observer
type Monitor2 struct{}

func (m *Monitor2) Name() string {
	return "monitor2"
}

func (m *Monitor2) Update(msg string) {
	fmt.Printf("monitor 2 updates %s\n", msg)
}

func Run() {
	central := Central{monitors: map[string]Observer{}}
	central.Register(&Monitor1{})
	central.Register(&Monitor2{})
	central.Notify("hi")
}
