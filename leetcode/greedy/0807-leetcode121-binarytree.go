package dynamicprogram

func maxProfit(prices []int) int {
	buy := prices[0]
	maxProfile := 0
	for _, sell := range prices {
		profile := sell - buy
		if profile < 0 {
			buy = sell
			continue
		}
		if profile > maxProfile {
			maxProfile = profile
		}
	}
	return maxProfile
}

// 超出时间限制
//func maxProfit(prices []int) int {
//	max := 0
//	n := len(prices)
//	for i := 0; i < n; i++ {
//		for j := i; j < n; j++ {
//			if prices[j] - prices[i] > max {
//				max = prices[j] - prices[i]
//			}
//		}
//	}
//	return max
//}
