package leetcode

func divisorGame(n int) bool {
	if n == 0 ||  n == 1 {
		return false
	}
	results := make([]bool, n+1)
	results[0], results[1] = true, true
	for i := 2; i <= n-1; i++ {
		for j := 1; j <= i/2; j++ {
			if n % j == 0 && results[i-j] == false {
				results[i] = false
				break
			}
		}
	}
	return results[n]
}
