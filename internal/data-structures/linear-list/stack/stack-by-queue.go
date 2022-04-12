package stack

// StackQ implemented using 2 queues
type StackQ struct {
	q1 *Queue
	q2 *Queue
}

func (s *StackQ) Push(o interface{}) {
	if s.q1.Size() != 0 {
		s.q1.Enqueue(o)
	} else {
		s.q2.Enqueue(o)
	}
}

func (s *StackQ) Pop() (o interface{}, err error) {
	if s.q1.Size() != 0 {
		return s.popFromNonEmptyQueue(s.q2, s.q1)
	}
	return s.popFromNonEmptyQueue(s.q1, s.q2)
}

func (s *StackQ) Top() (o interface{}, err error) {
	if s.q1.Size() != 0 {
		item, err := s.popFromNonEmptyQueue(s.q2, s.q1)
		if err != nil {
			return nil, err
		}
		s.q2.Enqueue(item)
		return item, nil
	}
	item, err := s.popFromNonEmptyQueue(s.q1, s.q2)
	if err != nil {
		return nil, err
	}
	s.q1.Enqueue(item)
	return item, nil
}

func (s *StackQ) Size() int {
	return s.q1.Size() + s.q2.Size()
}

func (s *StackQ) IsEmpty() bool {
	return s.q1.Size()+s.q2.Size() == 0
}

func (s *StackQ) popFromNonEmptyQueue(emptyQ, nonEmptyQ *Queue) (o interface{}, err error) {
	for nonEmptyQ.Size() > 1 {
		item, err := nonEmptyQ.Dequeue()
		if err != nil {
			return nil, err
		}
		emptyQ.Enqueue(item)
	}
	item, err := nonEmptyQ.Dequeue()
	if err != nil {
		return nil, err
	}
	return item, nil
}
