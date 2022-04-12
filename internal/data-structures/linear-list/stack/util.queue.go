package stack

import (
	"container/list"
	"errors"
)

// Queue implemented using linked list
type Queue struct {
	list *list.List
}

func (q *Queue) Enqueue(o interface{}) {
	q.list.PushBack(o)
}

func (q *Queue) Dequeue() (o interface{}, err error) {
	if q.Size() == 0 {
		return nil, errors.New("empty queue")
	}
	return q.list.Remove(q.list.Front()), nil
}

func (q *Queue) Front() (o interface{}, err error) {
	if q.Size() == 0 {
		return nil, errors.New("empty queue")
	}
	return q.list.Front().Value, nil
}

func (q *Queue) Size() int {
	return q.list.Len()
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}
