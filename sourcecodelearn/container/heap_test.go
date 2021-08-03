package container

import (
	"container/heap"
	"testing"
)

type IntSlice []int

func (s *IntSlice) Len() int {
	return len(*s)
}

func (s *IntSlice) Less(i, j int) bool {
	return (*s)[i] < (*s)[j]
}

func (s *IntSlice) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *IntSlice) Push(v interface{}) {
	*s = append(*s, v.(int))
}

func (s *IntSlice) Pop() interface{} {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func Test_IntSliceBasedHeap(t *testing.T) {
	println((-1) / 2)
	hp := IntSlice{4, 3, 2, 1}
	heap.Init(&hp)
	for _, i := range hp {
		t.Log(i)
	}
}

type Power struct {
	pow      int
	priority int
	index    int
}

type PowerList []Power

func (pl *PowerList) Len() int {
	return len(*pl)
}

func (pl *PowerList) Less(i, j int) bool {
	var power1, power2 Power
	for k, power := range *pl {
		if power.index == i {
			power1 = (*pl)[k]
		}
		if power.index == j {
			power2 = (*pl)[k]
		}
	}
	return power1.pow < power2.pow || (power1.pow == power2.pow && power1.priority > power2.priority)
}

func (pl *PowerList) Swap(i, j int) {
	var power1, power2 *Power
	for k, power := range *pl {
		if power.index == i {
			power1 = &(*pl)[k]
		}
		if power.index == j {
			power2 = &(*pl)[k]
		}
	}	
	(*power1).index, (*power2).index = j, i
}

func (pl *PowerList) Push(v interface{}) {
	panic("implement me")
}

func (pl *PowerList) Pop() interface{} {
	panic("implement me")
}
