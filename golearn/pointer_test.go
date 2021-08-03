package golearn

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPointType(t *testing.T) {
	var a int
	b := &a
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
}

func TestSliceArgs(t *testing.T) {
	a := []int{1,2,3,4,5}
	b := a[:]
	update(b, 2, 4)
	for _, i := range a {
		fmt.Println(i)
	}
}

func update(list []int, i int, val int) {
	list[i] = val
}