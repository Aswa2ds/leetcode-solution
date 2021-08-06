package golearn

import (
	"fmt"
	"testing"
)

func TestDeferAndPanic(t *testing.T) {
	deferAndPanic()
}

func deferAndPanic() {
	defer func() {fmt.Println("msg1")}()
	defer func() {fmt.Println("msg2")}()
	defer func() {fmt.Println("msg3")}()

	panic("msg4")
}
