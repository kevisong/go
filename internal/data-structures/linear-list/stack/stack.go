package stack

// IStack defines the stack interface
type IStack interface {
	Push(interface{})
	Pop() (interface{}, error)

	Top() (interface{}, error)
	Size() int
	IsEmpty() bool
}
