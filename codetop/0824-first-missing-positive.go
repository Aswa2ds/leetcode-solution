package codetop

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range nums {
		if nums[i] <= 0 {
			nums[i] = n+1
		}
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
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

	for i := range nums {
		if nums[i] > 0 {
			return i+1
		}
	}
	return n+1

}
