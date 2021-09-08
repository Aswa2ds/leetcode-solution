package dichotomy

func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return 0
}

// 可用hash O(n) 可解，但空间超了
// func findDuplicate(nums []int) int {
// 	m := make(map[int]int)
// 	for _, num := range nums {
// 		if m[num] == 1 {
// 			return num
// 		}
// 		m[num]++
// 	}
// 	return 0
// }