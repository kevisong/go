package deque

import (
	"container/list"
	"errors"
)

// IDeque defines the deque interface
type IDeque interface {
	InsertFront(o interface{})
	InsertRear(o interface{})
	DeleteFront() (o interface{}, err error)
	DeleteRear() (o interface{}, err error)

	GetFront() (o interface{}, err error)
	GetRear() (o interface{}, err error)
	Size() int
	IsEmpty() bool
}

type Deque struct {
	list *list.List
}

func (d *Deque) InsertFront(o interface{}) {
	d.list.PushFront(o)
}
func (d *Deque) InsertRear(o interface{}) {
	d.list.PushBack(o)
}
func (d *Deque) DeleteFront() (o interface{}, err error) {
	if d.list.Len() == 0 {
		return nil, errors.New("empty deque")
	}
	return d.list.Remove(d.list.Front()), err
}
func (d *Deque) DeleteRear() (o interface{}, err error) {
	if d.list.Len() == 0 {
		return nil, errors.New("empty deque")
	}
	return d.list.Remove(d.list.Back()), err
}

func (d *Deque) GetFront() (o interface{}, err error) {
	if d.list.Len() == 0 {
		return nil, errors.New("empty deque")
	}
	return d.list.Front().Value, err
}
func (d *Deque) GetRear() (o interface{}, err error) {
	if d.list.Len() == 0 {
		return nil, errors.New("empty deque")
	}
	return d.list.Back().Value, err
}
func (d *Deque) Size() int {
	return d.list.Len()
}
func (d *Deque) IsEmpty() bool {
	return d.list.Len() == 0
}
