package binarytree

type IHeap interface {
	// swap with parent if unordered
	MoveUp()
	// swap with child if unordered, aka. heapify
	MoveDown()

	// push at the end, then use MoveUp() to re-heap
	Push()
	// remove at the front, move last node to front, then use MoveDown() to re-heap
	Pop()

	// remove any node
	// swap it with last node
	// if it's child out of order, use MoveDown()
	// if it's parent out of order, use MoveUp()
	Remove(interface{})

	// keep calling Push() to turn an unordered array to a heap
	Build()
}

type Heap struct {
	array []int
}
