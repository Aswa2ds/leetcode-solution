package main

import (
	"fmt"
	"strings"
)

//func main() {
//	c := make(chan int)
//
//	go func() {
//		//c <- 1 // send to channel
//		close(c)
//	}()
//
//	x := <-c // recv from channel
//
//	fmt.Println(x)
//}

func main() {
	s := "    12"
	s = strings.TrimLeft(s, " ")
	fmt.Println(s)
}

//func foo() int {
//	a := 1
//	return a
//}

//func foo() *int {
//	a := new(int)
//	return a
//}

//func foo() {
//	var _ = make([]int, 100)
//	foo()
//}
