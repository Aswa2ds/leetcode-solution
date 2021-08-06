package hash

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i, num := range nums {
		if num <= 0 {
			nums[i] = n+1
		}
	}
	abs := func(x int) int {
		if x < 0 {
			x = -x
		}
		return x
	}
	for _, num := range nums {
		num = abs(num)
		if num > n {
			continue
		}
		if nums[num-1] > 0 {
			nums[num-1] = -nums[num-1]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i+1
		}
	}
	return n+1
}

//func firstMissingPositive(nums []int) int {
//	n := len(nums)
//	nums = append(nums, -1)
//	for i, num := range nums {
//		if num <= 0 || num > n {
//			nums[i] = -1
//		}
//	}
//	for i := 0; i <= n; i++ {
//		for nums[i] != -1 && i != nums[i] && nums[i] != nums[nums[i]] {
//			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
//		}
//		if nums[i] != -1 && nums[i] != i {
//			nums[i] = -1
//		}
//	}
//	for i := 1; i <= n; i++ {
//		if nums[i] == -1 {
//			return i
//		}
//	}
//	return n+1
//}
