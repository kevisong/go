package queue

import (
	"container/list"
	"errors"
)

// QueueL implemented using linked list
type QueueL struct {
	list *list.List
}

func (q *QueueL) Enqueue(o interface{}) {
	q.list.PushBack(o)
}

func (q *QueueL) Dequeue() (o interface{}, err error) {
	if q.Size() == 0 {
		return nil, errors.New("empty queue")
	}
	return q.list.Remove(q.list.Front()), nil
}

func (q *QueueL) Front() (o interface{}, err error) {
	if q.Size() == 0 {
		return nil, errors.New("empty queue")
	}
	return q.list.Front().Value, nil
}

func (q *QueueL) Size() int {
	return q.list.Len()
}

func (q *QueueL) IsEmpty() bool {
	return q.Size() == 0
}
