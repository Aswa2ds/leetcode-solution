package leetcode

import (
	"fmt"
	"testing"
)

func Test_divisorGame(t *testing.T) {
	for i := 0; i < 20; i++ {
		fmt.Println(divisorGame(i))
	}
}
