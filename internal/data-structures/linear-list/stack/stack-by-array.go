package stack

import "errors"

// StackA implemented using array
type StackA struct {
	array []interface{}
}

func (s *StackA) Push(o interface{}) {
	s.array = append(s.array, o)
}

func (s *StackA) Pop() (o interface{}, err error) {
	if s.Size() == 0 {
		return nil, errors.New("empty stack")
	}
	top := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return top, nil
}

func (s *StackA) Top() (o interface{}, err error) {
	if s.Size() == 0 {
		return nil, errors.New("empty stack")
	}
	return s.array[len(s.array)-1], nil
}

func (s *StackA) Size() int {
	return len(s.array)
}

func (s *StackA) IsEmpty() bool {
	return s.Size() == 0
}
