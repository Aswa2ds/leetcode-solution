package dynamicprogram

func canJump(nums []int) bool {
	furthestReachable := 0
	for i, num := range nums {
		if furthestReachable >= i {
			if furthestReachable < i + num {
				furthestReachable = i + num
			}
		} else {
			break
		}
	}
	return furthestReachable >= len(nums)-1
}
