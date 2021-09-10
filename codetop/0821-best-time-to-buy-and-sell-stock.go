package codetop

//func maxProfit(prices []int) int {
//	buy := prices[0]
//	profile := 0
//	for _, price := range prices {
//		tmp := price - buy
//		if tmp < 0 {
//			buy = price
//		} else {
//			if tmp > profile {
//				profile = tmp
//			}
//		}
//	}
//	return profile
//}

func maxProfiti(prices []int) int {
	dp := make([]int, 2)
	dp[0] = 0
	dp[1] = -prices[0]
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < len(prices); i++ {
		dp[0] = max(dp[0], dp[1]+prices[i])
		dp[1] = max(dp[1], -prices[i])
	}
	return dp[0]
}
