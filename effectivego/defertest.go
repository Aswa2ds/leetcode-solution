package effectivego

import "fmt"

func trace(s string) string {
	fmt.Printf("enter: %v\n", s)
	return s
}

func un(s string) {
	fmt.Printf("leave: %v\n", s)
}

func DeferTest() {
	defer un(trace("b"))
	fmt.Println("in b")
}