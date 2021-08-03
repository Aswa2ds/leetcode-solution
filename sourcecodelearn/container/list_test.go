package container

import (
	"container/list"
	"fmt"
	"testing"
)

func Test_ListOperation(t *testing.T) {
	l := list.New()
	l.Init()
	l.PushBack(1)
	l.PushFront(2)
	for p := l.Front(); p != nil; p = p.Next() {
		if p.Value.(int) == 1 {
			l.InsertAfter(3, p)
		}
		if p.Value.(int) == 2 {
			l.InsertBefore(4, p)
		}
	}
	for p := l.Front(); p != nil; p = p.Next() {
		fmt.Println(p.Value.(int))
	}
}
