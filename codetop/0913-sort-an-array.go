package codetop

import (
	"math/rand"
	"time"
)

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivotIndex := partition(nums, left, right)
	quickSort(nums, left, pivotIndex-1)
	quickSort(nums, pivotIndex+1, right)
}

func partition(nums []int, left, right int) int {
	rand.Seed(time.Now().UnixMicro())
	key := rand.Intn(right-left) + left
	pivot := nums[key]
	nums[left], nums[key] = nums[key], nums[left]
	lo, hi := left, right
	for lo < hi {
		for lo < hi && nums[hi] >= pivot {
			hi--
		}
		for lo < hi && nums[lo] <= pivot {
			lo++
		}
		nums[lo], nums[hi] = nums[hi], nums[lo]
	}
	nums[left], nums[lo] = nums[lo], nums[left]
	return lo
}

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}
