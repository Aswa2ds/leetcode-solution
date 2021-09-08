package golearn

import (
	"fmt"
)

func DeferAndPanic() {
	defer func() {fmt.Println("msg1")}()
	defer func() {fmt.Println("msg2")}()
	defer func() {fmt.Println("msg3")}()

	panic("msg4")
}
