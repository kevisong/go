package templatemethod

import "fmt"

// The intent of the template method is to define the overall structure
// of the operation, while allowing subclasses to refine, or redefine,
// certain steps.

type IBurger interface {
	addMeat()
	addVege()
	addCheese()
	wantMeat() bool
	wantVege() bool
	wantCheese() bool
}

// Burger is the base class that contains the template method - MakeBurger()
type Burger struct {
	IBurger
}

func (b *Burger) cutBun() {}
func (b *Burger) wrapUp() {}

func (b *Burger) MakeBurger() {
	b.cutBun()
	if b.wantMeat() {
		b.addMeat()
	}
	if b.wantVege() {
		b.addVege()
	}
	if b.wantCheese() {
		b.addCheese()
	}
	b.wrapUp()
}

// Cheeseburger is the subclass. Go does not have inheritence,
// instead, composition is used to access other object's method,
// in this case, Burger.MakeBurger().
type Cheeseburger struct {
	Burger
}

func (c *Cheeseburger) wantMeat() bool {
	return true
}
func (c *Cheeseburger) addMeat() {
	fmt.Println("add meat")
}
func (c *Cheeseburger) wantVege() bool {
	return false
}
func (c *Cheeseburger) addVege() {
	fmt.Println("add vege")
}
func (c *Cheeseburger) wantCheese() bool {
	return true
}
func (c *Cheeseburger) addCheese() {
	fmt.Println("add cheese")
}

func Run() {
	cheeseburger := Cheeseburger{}
	cheeseburger.Burger = Burger{
		IBurger: &cheeseburger,
	}
	cheeseburger.MakeBurger()
}
