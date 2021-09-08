package codetop

import (
	"math/rand"
	"time"
)

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return make([][]int, 0)
	}
	var quickSort func([]int, int, int)
	var partition func([]int, int, int) int
	quickSort = func(nums []int, left int, right int) {
		if left >= right {
			return
		}
		pivotIndex := partition(nums, left, right)
		quickSort(nums, left, pivotIndex-1)
		quickSort(nums, pivotIndex+1, right)
	}
	partition = func(nums []int, left int, right int) int {
		rand.Seed(time.Now().UnixNano())
		key := rand.Intn(right-left) + left
		nums[left], nums[key] = nums[key], nums[left]
		lo, hi, pivot := left, right, nums[left]
		for lo < hi {
			for lo < hi && nums[hi] >= pivot { hi-- }
			for lo < hi && nums[lo] <= pivot { lo++ }
			nums[lo], nums[hi] = nums[hi], nums[lo]
		}
		nums[left], nums[lo] = nums[lo], nums[left]
		return lo
	}
	quickSort(nums, 0, len(nums)-1)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < k && nums[i] + nums[j] + nums[k] > 0 {
				k--
			}
			if j == k {
				break
			}
			if nums[i] + nums[j] + nums[k] == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return res
}