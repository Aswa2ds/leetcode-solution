package codetop

import (
	"fmt"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	list := []int{1,3,2}
	nextPermutation(list)
	fmt.Println(list)
}
