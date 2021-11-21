package prototype

import "fmt"

// Rectangle is a prototype
type Rectangle struct {
	x     int
	y     int
	color string
}

// Clone returns a copy of the prototype
func (r Rectangle) Clone() Rectangle {
	return r
}

func Run() {
	r1 := Rectangle{x: 0, y: 0, color: "blue"}
	r2 := r1.Clone()
	fmt.Println(r1)
	fmt.Println(r2)
}
