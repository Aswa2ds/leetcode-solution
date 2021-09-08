package codetop

import (
	"math/rand"
	"time"
)

func merge(intervals [][]int) [][]int {
	var quickSort func(intervals [][]int, left, right int)
	var partition func(internals [][]int, left, right int) int
	quickSort = func(intervals [][]int, left, right int) {
		if left >= right {
			return
		}
		pivotIndex := partition(intervals, left, right)
		quickSort(intervals, left, pivotIndex-1)
		quickSort(intervals, pivotIndex+1, right)
	}
	partition = func(internals [][]int, left, right int) int {
		rand.Seed(time.Now().UnixNano())
		key := rand.Intn(right-left) + left
		internals[left], intervals[key] = intervals[key], intervals[left]
		pivot := intervals[left]
		lo, hi := left, right
		for lo < hi {
			for lo < hi && intervals[hi][0] >= pivot[0] {
				hi--
			}
			for lo < hi && intervals[lo][0] <= pivot[0] {
				lo++
			}
			intervals[lo], intervals[hi] = intervals[hi], intervals[lo]
		}
		intervals[lo], intervals[left] = intervals[left], intervals[lo]
		return lo
	}
	quickSort(intervals, 0, len(intervals)-1)
	ret := make([][]int, 0)
	index := -1
	for _, interval := range intervals {
		if index < 0 || interval[0] > ret[index][1] {
			ret = append(ret, interval)
			index++
		} else {
			if interval[1] > ret[index][1] {
				ret[index][1] = interval[1]
			}
		}
	}
	return ret
}
