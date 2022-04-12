package queue

import "errors"

// QueueS implemented using 2 stacks
type QueueS struct {
	stackIn  *Stack
	stackOut *Stack
}

func (q *QueueS) Enqueue(o interface{}) {
	q.stackIn.Push(o)
}

func (q *QueueS) Dequeue() (o interface{}, err error) {
	if err := q.shift(); err != nil {
		return nil, err
	}
	if q.stackOut.Size() == 0 {
		return nil, errors.New("empty queue")
	}
	return q.stackOut.Pop()
}

func (q *QueueS) Front() (o interface{}, err error) {
	if err := q.shift(); err != nil {
		return nil, err
	}
	return q.stackOut.Top()
}

func (q *QueueS) Size() int {
	return q.stackIn.Size() + q.stackOut.Size()
}

func (q *QueueS) IsEmpty() bool {
	return q.Size() == 0
}

func (q *QueueS) shift() error {
	for !q.stackIn.IsEmpty() {
		item, err := q.stackIn.Pop()
		if err != nil {
			return err
		}
		q.stackOut.Push(item)
	}
	return nil
}
