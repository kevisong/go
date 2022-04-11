package linkedlist

// IList defines the linked list interface
type IList interface {
	// read operations
	Len() int
	Front() *Node
	Back() *Node

	// write operations
	PushFront(*Node)
	PushBack(*Node)
	InsertBefore(n, target *Node)
	InsertAfter(n, target *Node)

	// delete operation
	Remove(*Node)

	// advance operations
	Reverse()
}
