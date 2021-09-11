package codetop

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m + n - 1
	j, k := m-1, n-1
	for j >= 0 && k >= 0 {
		if nums1[j] >= nums2[k] {
			nums1[i] = nums1[j]
			j--
			i--
		} else {
			nums1[i] = nums2[k]
			k--
			i--
		}
	}
	for k >= 0 {
		nums1[i] = nums2[k]
		i--
		k--
	}
}
