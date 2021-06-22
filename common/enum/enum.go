package enum

// Direction defines direction
type Direction int

const (
	North Direction = iota
	East
	South
	West
)

// String returns string representation of Direction
func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}
