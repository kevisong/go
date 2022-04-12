package queue

import "errors"

// QueueA implemented using array
type QueueA struct {
	array []interface{}
}

func (q *QueueA) Enqueue(o interface{}) {
	q.array = append(q.array, o)
}

func (q *QueueA) Dequeue() (o interface{}, err error) {
	if q.Size() == 0 {
		return nil, errors.New("empty queue")
	}
	front := q.array[0]
	q.array = q.array[1:]
	return front, nil
}

func (q *QueueA) Front() (o interface{}, err error) {
	if q.Size() == 0 {
		return nil, errors.New("empty queue")
	}
	return q.array[0], nil
}

func (q *QueueA) Size() int {
	return len(q.array)
}

func (q *QueueA) IsEmpty() bool {
	return q.Size() == 0
}
