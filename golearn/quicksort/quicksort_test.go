package quicksort

import (
	"fmt"
	"testing"
)

type IntSlice []int

func (s *IntSlice) Compare(i int, j int) bool {
	return (*s)[i] < (*s)[j]
}

func (s *IntSlice) Swap(i int, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *IntSlice) Len() int {
	return len(*s)
}

func TestSort(t *testing.T) {
	list := IntSlice{2,45,2,354,23,12,56,0,54,3,5,67,876,43,23,5,789,87,1}
	Sort(&list)
	fmt.Println(list)
}
