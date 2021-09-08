package codetop

func longestValidParentheses(s string) int {
	var ret int
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] > 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		ret = max(dp[i], ret)
	}
	return ret
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

/* var ret int
   dp := make([]int, len(s))
   for i := 1; i < len(s); i++ {
   	if s[i] == ')' {
   		if s[i-1] == '(' {
   			if i >= 2 {
   				dp[i] = dp[i-2] + 2
   			} else {
   				dp[i] = 2
   			}
   		} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
   			if i-dp[i-1] >= 2 {
   				dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
   			} else {
   				dp[i] = dp[i-1] + 2
   			}
   		}
   	}
   	ret = max(ret, dp[i])
   }
   return ret */
