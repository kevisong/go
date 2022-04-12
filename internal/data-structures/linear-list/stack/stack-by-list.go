package stack

import (
	"container/list"
	"errors"
)

// StackL implemented using linked list
type StackL struct {
	list *list.List
}

func (s *StackL) Push(o interface{}) {
	s.list.PushBack(o)
}

func (s *StackL) Pop() (o interface{}, err error) {
	if s.list.Len() == 0 {
		return nil, errors.New("empty stack")
	}
	return s.list.Remove(s.list.Back()), nil
}

func (s *StackL) Top() (o interface{}, err error) {
	if s.list.Len() == 0 {
		return nil, errors.New("empty stack")
	}
	return s.list.Back().Value, nil
}

func (s *StackL) Size() int {
	return s.list.Len()
}

func (s *StackL) IsEmpty() bool {
	return s.list.Len() == 0
}
