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

func TestEmptyString(t *testing.T) {
	str := ""
	fmt.Println(str)
}

func TestSliceType(t *testing.T)  {
	a := []int{1,2,3,4}
	b := make([]int, 4)
	for i := range b {
		b[i] = i+1
	}
	c := a[:]
	d := b[:]
	update(a, 2, 4)
	for _, i := range a {
		fmt.Println(i)
	}
	update(b, 2, 4)
	for _, i := range b {
		fmt.Println(i)
	}
	update(a, 2, 4)
	for _, i := range a {
		fmt.Println(i)
	}
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(d))
}