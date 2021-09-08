package slidewindow

func merge(nums1 []int, m int, nums2 []int, n int)  {
	index := m+n-1
	s1, s2 := m-1, n-1
	for s1 >= 0 && s2 >= 0 {
		if nums1[s1] > nums2[s2] {
			nums1[index] = nums1[s1]
			s1--
		} else {
			nums1[index] = nums2[s2]
			s2--
		}
		index--
	}
	for s2 >= 0 {
		nums1[index] = nums2[s2]
		index--
		s2--
	}
}
