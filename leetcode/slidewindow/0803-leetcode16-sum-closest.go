package slidewindow

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	res := math.MaxInt32
	sort.Ints(nums)
	for _, num := range nums {
		fmt.Println(num)
	}
	abs := func(x int) int {
		if x < 0 {
			x = -x
		}
		return x
	}
	for i := range nums {
		for j, k := i+1, len(nums)-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return target
			}
			if abs(sum - target) < abs(res - target) {
				res = sum
			}
			if sum < target {
				j++
			} else {
				k--
			}
		}
	}
	return res
}
