package linkedlist

// TODO: implement singly linked list

// Node defines the node in a linked list
type Node struct {
	Data interface{}
	next *Node
}

func (n *Node) Next() *Node {
	return n.next
}

type List struct {
	head *Node
	tail *Node
	len  int
}

// Init -
func (l *List) Init() *List {
	return l
}

// New returns an initialized doubly linked list.
func New() *List { return new(List).Init() }

func (l *List) Front() *Node {
	return l.head
}

func (l *List) Back() *Node {
	return l.tail
}

func (l *List) PushFront(n *Node) {
	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		n.next = l.head
		l.head = n
	}
}

func (l *List) PushBack(n *Node) {
	if l.tail == nil {
		l.head = n
		l.tail = n
	} else {
		l.tail.next = n
		l.tail = n
	}
}

func (l *List) InsertBefore(n, target *Node) {
	n.next = target
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Next() == target {
			e.next = n
		}
	}
}

func (l *List) InsertAfter(n, target *Node) {
	n.next = target.next
	target.next = n
}

func (l *List) Reverse() {
	var prev *Node
	e := l.Front()
	for e != nil {
		next := e.next // save next
		e.next = prev  // set e's next to prev
		prev = e       // save e as prev
		e = next       // move e forward
	}
	l.head = prev
}
