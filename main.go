package main

import (
	"aswa2ds.cn/leetcode-solution/offer"
	"fmt"
)

func main() {
	medianFinder := offer.Constructor()
	medianFinder.AddNum(-1)
	fmt.Println(medianFinder.FindMedian())
	medianFinder.AddNum(-2)
	fmt.Println(medianFinder.FindMedian())
	medianFinder.AddNum(-3)
	fmt.Println(medianFinder.FindMedian())
	medianFinder.AddNum(-4)
	fmt.Println(medianFinder.FindMedian())
	medianFinder.AddNum(-5)
	fmt.Println(medianFinder.FindMedian())
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