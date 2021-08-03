package effectivego

import "fmt"

func ShiftAndAdd() {
	// 利用fmt 表示运算关系，而不是通过空格
	result1 := 1<<2 + 3<<2
	result2 := 1<<2 + 3<<2
	fmt.Println(result1)
	fmt.Println(result2)
}
