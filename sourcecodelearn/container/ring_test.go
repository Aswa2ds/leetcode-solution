package container

import (
	"container/ring"
	"fmt"
	"testing"
)

func Test_RingOperation(t *testing.T) {
	r := ring.New(5)
	for p, i := r, 0; i < r.Len(); p, i = p.Next(), i+1 {
		p.Value = i
		fmt.Println(p.Value)
	}
	for true {
		r = r.Move(1)
		if r.Value == 3 {
			fmt.Println("move to 3")
			break
		}
	}
	r2 := ring.New(5)
	for p, i := r2, 0; i < r2.Len(); p, i = p.Next(), i+1 {
		p.Value = 100
	}
	r.Link(r2)
	for i := 0; i < r.Len(); i++ {
		r = r.Move(1)
		fmt.Println(r.Value)
	}
}
