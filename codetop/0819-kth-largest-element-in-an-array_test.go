package codetop

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				nums: []int{5,2,4,1,3,6,0},
				k:    4,
			},
			want: 3,
		},

		{
			name: "test",
			args: args{
				nums: []int{1},
				k:    1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quickSort(t *testing.T) {
	list := generateRandomList(1000)
	quickSort(list, 0, len(list)-1)
	fmt.Println(list)
}

func generateRandomList(n int) []int {
	list := make([]int, n)
	rand.Seed(time.Now().UnixNano())
	for i := range list {
		list[i] = rand.Intn(1000000)
	}
	return list
}