package mathskill

func countDigitOne(n int) int {
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
	ret := 0
	for i := 1; i <= n; i *= 10 {
		driver := i*10
		ret += n/driver * i + min(max(n%driver-i+1, 0), i)
	}
	return ret
}
