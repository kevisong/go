package linkedlist

// implement doubly linked list

// NodeD defines the node in a doubly linked list
type NodeD struct {
	Data interface{}
	prev *NodeD
	next *NodeD
}

func (n *NodeD) Next() *NodeD {
	return n.next
}

func (n *NodeD) Prev() *NodeD {
	return n.prev
}

// ListD defines a doubly linked list
// n <---> ... <--> n
type ListD struct {
	head *NodeD
	tail *NodeD
	len  int
}

func (l *ListD) Front() *NodeD {
	if l.len == 0 {
		return nil
	}
	return l.head
}

func (l *ListD) Back() *NodeD {
	if l.len == 0 {
		return nil
	}
	return l.tail
}

// PushFront inserts the node before head
func (l *ListD) PushFront(n *NodeD) {
	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		n.next = l.head
		n.next.prev = n
		l.head = n
	}
	l.len++
}

// PushBack inserts the node after tail
func (l *ListD) PushBack(n *NodeD) {
	if l.tail == nil {
		l.tail = n
		l.head = n
	} else {
		n.prev = l.tail
		n.prev.next = n
		l.tail = n
	}
	l.len++
}

// InsertBefore inserts the node before target
func (l *ListD) InsertBefore(n, target *NodeD) {
	n.prev = target.prev
	n.next = target
	n.prev.next = n
	n.next.prev = n
	l.len++
}

// InsertAfter inserts the node after target
func (l *ListD) InsertAfter(n, target *NodeD) {
	n.prev = target
	n.next = target.next
	n.prev.next = n
	n.next.prev = n
	l.len++
}

// Remove deletes the node
func (l *ListD) Remove(n *NodeD) {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.prev = nil
	n.next = nil
	l.len--
}

// Reverse reverse the list
func (l *ListD) Reverse() {
	// revsersing head and tail
	l.head, l.tail = l.tail, l.head
	// reversing nodes
	for n := l.Front(); n != nil; n = n.Next() {
		temp := n.prev
		n.prev = n.next
		n.next = temp
	}
}

// Init -
func (l *ListD) Init() *ListD {
	return l
}

// NewD returns an initialized doubly linked list.
func NewD() *ListD { return new(ListD).Init() }

func (l *ListD) Len() int { return l.len }
