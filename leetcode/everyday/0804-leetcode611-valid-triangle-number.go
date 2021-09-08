package everyday

import "sort"

func triangleNumber(nums []int) (ans int) {
	sort.Ints(nums)
	for i, num := range nums {
		for j := i+1; j < len(nums); j++ {
			tmp := search(nums[j+1:], num+nums[j])
			ans += tmp
		}
	}
	return ans
}

func search(nums []int, v int) int {
	low, high := 0, len(nums)
	for low < high {
		mid := (low + high) >> 1
		if nums[mid] < v {
			low = mid+1
		} else {
			high = mid
		}
	}
	return low
}
