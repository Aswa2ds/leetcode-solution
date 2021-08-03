package offer

func majorityElement(nums []int) int {
	card, point := 0, 0
	for _, num := range nums {
		if point == 0 {
			card = num
		}
		if num == card {
			point++
		} else {
			point--
		}
	}
	return card
}
