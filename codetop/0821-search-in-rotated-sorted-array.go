package codetop

func search(nums []int, target int) int {
	var partition func([]int, int, int, int) int
	partition = func(nums []int, target int, lo int, hi int) int {
		for lo < hi {
			mid := (lo + hi) >> 1
			pivot := nums[mid]
			if nums[lo] <= pivot {
				if target < nums[lo] || target > pivot {
					lo = mid + 1
				} else {
					hi = mid
				}
			} else {
				if target < nums[lo] && target > pivot {
					lo = mid+1
				} else {
					hi = mid
				}
			}
		}
		return lo
	}
	ret := partition(nums, target, 0, len(nums))
	if ret >= 0 && ret < len(nums) && nums[ret] == target {
		return ret
	}
	return -1
}

//func search(nums []int, v int) int {
//	low, high := 0, len(nums)
//	for low < high {
//		mid := (low + high) >> 1
//		if nums[mid] < v {
//			low = mid+1
//		} else {
//			high = mid
//		}
//	}
//	return low
//}