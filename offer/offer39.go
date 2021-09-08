package offer

//func majorityElement(nums []int) int {
//	card, point := 0, 0
//	for _, num := range nums {
//		if point == 0 {
//			card = num
//		}
//		if num == card {
//			point++
//		} else {
//			point--
//		}
//	}
//	return card
//}

func mooreVote(nums []int) int {
	card := 0
	candidate := nums[1]
	for _, num := range nums {
		if num == candidate {
			card++
		} else {
			card--
		}
		if card == 0 {
			candidate = num
		}
	}
	return candidate
}
