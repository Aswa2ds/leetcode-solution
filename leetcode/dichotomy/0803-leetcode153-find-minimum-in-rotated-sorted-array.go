package dichotomy

func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := (low+high)/2
		if nums[mid] < nums[high] {
			high = mid
			continue
		}
		if nums[mid] > nums[high] {
			low = mid+1
			continue
		}
		if nums[mid] == nums[low] {
			high--
		}
	}
	return nums[low]
}