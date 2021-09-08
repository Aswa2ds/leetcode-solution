package slice

import (
	"fmt"
	"testing"
)

func TestSliceScale(t *testing.T) {
	list := make([]int, 0)
	oldCap := cap(list)
	for i := 0; i < 2048; i++ {
		list = append(list, i)
		newCap := cap(list)
		if oldCap != newCap {
			fmt.Printf("%d --> %d\n", oldCap, newCap)
			oldCap = newCap
		}
	}
}
