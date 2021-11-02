package iterator

import "fmt"

type Iterator interface {
	HasNext() bool
	Next()
	CurrentItem() interface{}
}

type Player struct{}

type ArrayPlayer []Player

func (a ArrayPlayer) Iterator() Iterator {
	return &ArrayPlayerIterator{
		arrayPlayer: a,
		index:       0,
	}
}

type ArrayPlayerIterator struct {
	arrayPlayer ArrayPlayer
	index       int
}

func (a *ArrayPlayerIterator) HasNext() bool {
	return a.index < len(a.arrayPlayer)-1
}

func (a *ArrayPlayerIterator) Next() {
	a.index++
}

func (a *ArrayPlayerIterator) CurrentItem() interface{} {
	return a.arrayPlayer[a.index]
}

func Run() {
	arrayPlayer := ArrayPlayer{Player{}, Player{}}
	iterator := arrayPlayer.Iterator()
	for iterator.HasNext() {
		fmt.Println(iterator.CurrentItem())
		iterator.Next()
	}
}
