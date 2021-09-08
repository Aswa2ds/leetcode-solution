package codetop

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	} 
	_, n := len(matrix), len(matrix[0])
	dp := make([]int, n+1)
	//for i := range dp {
	//	dp[i] = make([]int, n+1)
	//}
	min := func(a, b int) int {
	    if a < b {
	        return a
	    }
	    return b
	}
	max := func(a, b int) int {
	    if a > b {
	        return a
	    }
	    return b
	}
	maxside := 0
	for _, line := range matrix {
		tmp := dp[0]
		for i, b := range line {
			if b == '1' {
				tmp = min(dp[i], min(dp[i+1], tmp))+1
				tmp, dp[i+1] = dp[i+1], tmp
				maxside = max(dp[i+1], maxside)
			} else {
				tmp = 0
				tmp, dp[i+1] = dp[i+1], tmp
			}
		}
	}
	return maxside*maxside
	
}
