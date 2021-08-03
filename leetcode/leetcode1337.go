package leetcode

import "container/heap"

type pair struct {
	pow int
	index int
}

type hp []pair

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i].pow < h[j].pow || (h[i].pow == h[j].pow && h[i].index < h[j].index)
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(pair))
}

func (h *hp) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

func kWeakestRows(mat [][]int, k int) []int {
	h := hp{}
	for i, line := range mat {
		power := getLinePower(line)
		h = append(h, pair{
			pow:   power,
			index: i,
		})
	}
	heap.Init(&h)
	ans := make([]int, k)
	for i := range ans {
		ans[i] = heap.Pop(&h).(pair).index
	}
	return ans
}

func getLinePower(line []int) int {
	if line[0] == 0 {
		return 0
	}
	if line[len(line)-1] == 1 {
		return len(line)
	}
	low, high := 0, len(line)
	for low < high {
		mid := (low + high)/2
		if line[mid] == 0 {
			high = mid
		} else {
			low = mid+1
		}
	}
	return low
}
