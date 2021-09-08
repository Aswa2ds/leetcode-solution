package main

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

func add(a, b int) int {
	return a + b
}

func main() {
	sum := add(1, 2)
	sum++
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
//
