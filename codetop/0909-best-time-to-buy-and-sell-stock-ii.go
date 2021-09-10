package codetop

func maxProfit(prices []int) int {
	maxP := 0
	if len(prices) == 0 {
		return maxP
	}
	buy := prices[0]
	for _, price := range prices {
		tmp := price - buy
		if tmp > 0 {
			maxP += tmp
		}
		buy = price
	}
	return maxP
}
