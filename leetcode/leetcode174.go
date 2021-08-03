package leetcode

import "math"

func calculateMinimumHP(dungeon [][]int) int {
	if dungeon == nil || len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 1
	}
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[m-1][n] = 1
	dp[m][n-1] = 1
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	for i := m-1; i >= 0; i-- {
		for j := n-1; j >= 0; j-- {
			dp[i][j] = max(min(dp[i+1][j], dp[i][j+1])-dungeon[i][j], 1)
		}
	}
	return dp[0][0]
}

//func calculateMinimumHP(dungeon [][]int) int {
//	if dungeon == nil {
//		return 0
//	}
//	if len(dungeon) == 0 {
//		return 0
//	}
//	if len(dungeon[0]) == 0 {
//		return 0
//	}
//	n := len(dungeon[0])
//	leftHp, upHp := 0, make([]int, n)
//	leftMinHp, upMinHp := 0, make([]int, n)
//
//	min := func(a, b int) int {
//		if a > b {
//			return b
//		}
//		return a
//	}
//
//	//max := func(a, b int) int {
//	//	if a > b {
//	//		return a
//	//	}
//	//	return b
//	//}
//
//	//for i := range upHp {
//	//	upHp[i] = math.MinInt32
//	//}
//	//for i := range upMinHp {
//	//	upMinHp[i] = math.MinInt32
//	//}
//
//	for num, line := range dungeon {
//		//if num == 0 {
//		//	for i := range line {
//		//		upHp[i] = leftHp + line[i]
//		//		upMinHp[i] = min(leftMinHp, upHp[i])
//		//		leftHp = upHp[i]
//		//		leftMinHp = upMinHp[i]
//		//	}
//		//	continue
//		//}
//		for i := range line {
//			if num != 0 && (leftHp-leftMinHp) <= (upHp[i]-upMinHp[i]) {
//				upHp[i] = upHp[i] + line[i]
//				upMinHp[i] = min(upMinHp[i], upHp[i])
//			} else {
//				upHp[i] = leftHp + line[i]
//				upMinHp[i] = min(leftMinHp, upHp[i])
//			}
//			leftHp = upHp[i]
//			leftMinHp = upMinHp[i]
//		}
//		leftHp = math.MinInt32
//		leftMinHp = math.MinInt32
//	}
//	finalHp := upMinHp[n-1]
//	if finalHp <= 0 {
//		return 1-finalHp
//	}
//	return 1
//}