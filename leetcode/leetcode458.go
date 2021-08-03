package leetcode

import "math"

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	n := minutesToTest/minutesToDie + 1
	// consider n == 0
	return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(n))))
}
