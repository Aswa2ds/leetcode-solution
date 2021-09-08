package bitcalculate

func singleNumber(nums []int) int {
	once, twice := 0, 0
	for _, num := range nums {
		once = (once^num)&(^twice)
		twice = (twice^num)&(^once)
	}
	return once
}

// bit 按位统计
//func singleNumber(nums []int) int {
//
//	var count [64]int
//	for _, num := range nums {
//		for i := 0; i < 64; i++ {
//			count[i] += num & 0x1
//			num >>= 1
//		}
//	}
//	var res int
//	for i, c := range count {
//		if c%3 != 0 {
//			res |= 1<<i
//		}
//	}
//	return res
//}

// hash
// 空间复杂度 O(n), 时间复杂度O(n)
//func singleNumber(nums []int) int {
//	m := make(map[int]int)
//	for _, num := range nums {
//		_, ok := m[num]
//		if !ok {
//			m[num] = 0
//		}
//		m[num]++
//	}
//	for k, v := range m {
//		if v == 1 {
//			return k
//		}
//	}
//	return 0
//}

// sort & find
// 空间复杂度 O(1), 时间复杂度O(nlogn)
//func singleNumber(nums []int) int {
//	sort.Ints(nums)
//	for i := 0; i < len(nums)-1; {
//		if nums[i+1] == nums[i] {
//			i += 3
//		} else {
//			return nums[i]
//		}
//	}
//	return nums[len(nums)-1]
//}