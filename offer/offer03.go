package offer

//func findRepeatNumber(nums []int) int {
//	var list [100000]bool
//	for _, num := range nums {
//		if list[num] {
//			return num
//		}
//		list[num] = true
//	}
//	return 0
//}

//func findRepeatNumber(nums []int) int {
//	m := make(map[int]bool)
//	for _, num := range nums {
//		_, ok := m[num]
//		if ok {
//			return num
//		}
//		m[num] = true
//	}
//	return 0
//}

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num == i {
			continue
		}
		if nums[num] == num {
			return num
		} else {
			nums[i], nums[num] = nums[num], num
			i--
		}
	}
	return 0
}
