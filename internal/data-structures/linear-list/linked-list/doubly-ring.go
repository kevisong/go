package linkedlist

// implement doubly linked list using ring (resembles container/list)

// NodeDR defines the node in a doubly linked list
type NodeDR struct {
	Data interface{}
	prev *NodeDR
	next *NodeDR

	list *ListDR
}

func (n *NodeDR) Next() *NodeDR {
	if n.next == n.list.root {
		return nil
	}
	return n.next
}

func (n *NodeDR) Prev() *NodeDR {
	if n.prev == n.list.root {
		return nil
	}
	return n.prev
}

// ListDR defines a doubly linked list implemented using ring
//  |-------------------------|
// root <--> n <---> ... <--> n
type ListDR struct {
	root *NodeDR
	len  int
}

func (l *ListDR) Front() *NodeDR {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func (l *ListDR) Back() *NodeDR {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// PushFront inserts the node after root
// root <-x-> n <---> n <---> ...
//  |        |
//  |-n_new-|
func (l *ListDR) PushFront(n *NodeDR) {
	n.prev = l.root
	n.next = l.root.next
	n.prev.next = n
	n.next.prev = n
	n.list = l
	l.len++
}

// PushBack inserts the node before root
//	|-----------------------------|
//  |---------x----------|		  |
// root <---> ... <---> n <---> n_new
func (l *ListDR) PushBack(n *NodeDR) {
	n.prev = l.root.prev
	n.next = l.root
	n.prev.next = n
	n.next.prev = n
	n.list = l
	l.len++
}

// InsertBefore inserts the node before target
//			   |-n_new-|
//			   |       |
// ... n <---> n <-x-> n_target <---> ...
func (l *ListDR) InsertBefore(n, target *NodeDR) {
	n.prev = target.prev
	n.next = target
	n.prev.next = n
	n.next.prev = n
	n.list = l
	l.len++
}

// InsertAfter inserts the node after target
//			          |-n_new-|
//			          |       |
// ... n <---> n_target <-x-> n <---> ...
func (l *ListDR) InsertAfter(n, target *NodeDR) {
	n.prev = target
	n.next = target.next
	n.prev.next = n
	n.next.prev = n
	n.list = l
	l.len++
}

// Remove deletes the node
//     |---------------|
//     |               |
// ... n <-x-> n <-x-> n ...
func (l *ListDR) Remove(n *NodeDR) {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.prev = nil
	n.next = nil
	n.list = nil
	l.len--
}

// Reverse reverse the list
func (l *ListDR) Reverse() {
	// reversing root
	temp := l.root.prev
	l.root.prev = l.root.next
	l.root.next = temp
	// reversing nodes
	for n := l.Front(); n != nil; n = n.Next() {
		temp := n.prev
		n.prev = n.next
		n.next = temp
	}
}

func (l *ListDR) Init() *ListDR {
	l.root = &NodeDR{}
	l.root.prev = l.root
	l.root.next = l.root
	l.len = 0
	return l
}

// NewDR returns an initialized doubly linked ring.
func NewDR() *ListDR { return new(ListDR).Init() }

func (l *ListDR) Len() int { return l.len }
