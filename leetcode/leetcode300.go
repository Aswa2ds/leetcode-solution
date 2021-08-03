package leetcode

func lengthOfLIS(nums []int) int {
	var d []int
	d = append(d, nums[0])
	l := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= d[l] {
			// 可优化二分法
			//j := l
			//for ; j > 0; j-- {
			//	if d[j-1] < nums[i] {
			//		break
			//	}
			//}
			//d[j] = nums[i]
			low, high := 0, l
			for low < high {
				mid := (low + high) / 2
				switch {
				case d[mid] == nums[i]:
					low, high = mid, mid
				case d[mid] > nums[i]:
					high = mid
				default:
					low = mid+1
				}
			}
			d[low] = nums[i]
		} else {
			d = append(d, nums[i])
			l++
		}
	}
	return len(d)
}

// dp
//func lengthOfLIS(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//	count := make([]int, len(nums))
//	for i := range nums {
//		j := 0
//		max := 1
//		for ; j < i; j++ {
//			if nums[i] > nums[j] {
//				if count[j] >= max {
//					max = count[j] + 1
//				}
//			}
//		}
//		count[i] = max
//	}
//	max := count[0]
//	for _, c := range count {
//		if c > max {
//			max = c
//		}
//	}
//	return max
//}
