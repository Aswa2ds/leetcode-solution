package effectivego

import "fmt"

func IfInitTest() {
	a := 1
	if true {
		a := 2
		fmt.Println(a)
	}
	for true {
		fmt.Println("xxx")
	}
	fmt.Println(a)
}
