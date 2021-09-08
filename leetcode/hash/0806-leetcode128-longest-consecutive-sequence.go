package hash

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}
	length := 0
	for key := range m {
		_, ok := m[key-1]
		if ok {
			continue
		}
		s := key+1
		for m[s] {
			s++
		}
		if length < s-key {
			length = s-key
		}
	}
	return length
}

// 时间复杂度：O(nlong)
//func longestConsecutive(nums []int) int {
//	sort.Ints(nums)
//	s1 := 0
//	length := 0
//	for s1 < len(nums) {
//		tmp := 1
//		s2, s3 := s1, s1
//		for s3+1 < len(nums) {
//			if nums[s3+1] == nums[s3] {
//				s3++
//			}
//			if nums[s3+1] == nums[s3]+1 {
//				s3++
//				tmp++
//				s2 = s3
//			} else {
//				break
//			}
//		}
//		if length < tmp {
//			length = tmp
//		}
//		s1, s2 = s2+1, s2+1
//	}
//	return length
//}
