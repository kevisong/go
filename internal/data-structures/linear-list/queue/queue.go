package queue

// IQueue defines the queue interface
type IQueue interface {
	Enqueue(interface{})
	Dequeue() (interface{}, error)

	Front() (interface{}, error)
	Size() int
	IsEmpty() bool
}
