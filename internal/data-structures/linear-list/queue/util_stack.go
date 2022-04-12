package queue

import (
	"container/list"
	"errors"
)

// Stack implemented using linked list
type Stack struct {
	list *list.List
}

func (s *Stack) Push(o interface{}) {
	s.list.PushBack(o)
}

func (s *Stack) Pop() (o interface{}, err error) {
	if s.list.Len() == 0 {
		return nil, errors.New("empty stack")
	}
	return s.list.Remove(s.list.Back()), nil
}

func (s *Stack) Top() (o interface{}, err error) {
	if s.list.Len() == 0 {
		return nil, errors.New("empty stack")
	}
	return s.list.Back().Value, nil
}

func (s *Stack) Size() int {
	return s.list.Len()
}

func (s *Stack) IsEmpty() bool {
	return s.list.Len() == 0
}
