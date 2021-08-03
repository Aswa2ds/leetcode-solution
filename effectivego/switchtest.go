package effectivego

import "fmt"

func SwitchTest() {
	switch {
	case 1 > 2:
		fmt.Println("here")
	case 1 < 2:
		fmt.Println("there")
	}
}
