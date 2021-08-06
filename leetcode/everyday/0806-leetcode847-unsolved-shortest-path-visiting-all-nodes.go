package everyday

import (
	"fmt"
	"math"
)

func shortestPathLength(graph [][]int) int {
	min := math.MaxInt32
	for v := range graph {
		states := make([]int, len(graph))
		length := calculatePathLength(graph, states, v, -1)
		if length < min {
			min = length
			fmt.Println(v)
		}
	}
	return min
}

func calculatePathLength(graph [][]int, states []int, v int, src int) int {
	states[v] = 1
	length := 0
	for i, e := range graph[v] {
		if i == len(graph[v])-1 {
			states[v] = 2
		}
		if states[e] != 2 {
			length++
			length += calculatePathLength(graph, states, e, v)
		}
	}
	return length
}
