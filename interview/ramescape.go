package interview

import "fmt"

type A struct {
	s string
}

func foo(s string) *A {
	a := new(A)
	a.s = s
	return a
}

func RamEscape() {
	a := foo("hello")
	b := a.s + " world"
	c := b + "!"
	fmt.Println(c)
}
