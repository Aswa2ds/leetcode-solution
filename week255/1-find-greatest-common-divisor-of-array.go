package week255

import (
	"math"
)

func findGCD(nums []int) int {
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
	minValue, maxValue := math.MaxInt64, math.MinInt64
	for _, num := range nums {
		minValue = min(minValue, num)
		maxValue = max(maxValue, num)
	}
	for i := minValue; i >= 1; i-- {
		if minValue%i == 0 && maxValue%i == 0 {
			return i
		}
	}

	return 1
}
