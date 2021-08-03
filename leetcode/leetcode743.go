package leetcode

import "container/list"

// 这个题 还是邻接表 效率更高
func networkDelayTime(times [][]int, n int, k int) int {
	// prepare
	costs := make([]int, n+1)
	for i := range costs {
		costs[i] = -1
	}
	matrix := make([][]int, n+1)
	for i := range matrix {
		matrix[i] = make([]int, n+1)
		for j := range matrix[i] {
			matrix[i][j] = -1
		}
	}
	for _, time := range times {
		s, d, c := time[0], time[1], time[2]
		matrix[s][d] = c
	}

	//destCostMap := make(map[int]destCosts, len(times))
	//for _, time := range times {
	//	s, d, c := time[0], time[1], time[2]
	//	if _, ok := destCostMap[s]; !ok {
	//		destCostMap[s] = make(destCosts, 0)
	//	}
	//	destCostMap[s] = append(destCostMap[s], destCost{
	//		dest: d,
	//		cost: c,
	//	})
	//}
	l := list.New()
	l.PushBack(k)
	costs[k] = 0
	for l.Len() > 0 {
		curr := l.Remove(l.Front()).(int)
		dcs := matrix[curr]
		for i := 1; i <= n; i++ {
			if dcs[i] != -1 {
				nc := costs[curr] + dcs[i]
				if costs[i] == -1 || nc < costs[i] {
					costs[i] = nc
					l.PushBack(i)
				}
			}
		}
		//dcs, ok := destCostMap[curr]
		//if !ok {
		//	continue
		//}
		//for _, dc := range dcs {
		//	nc := costs[curr] + dc.cost
		//	if costs[dc.dest] == -1 || nc < costs[dc.dest] {
		//		costs[dc.dest] = nc
		//		l.PushBack(dc.dest)
		//	}
		//}
	}
	maxCost := 0
	for i := 1; i < len(costs); i++ {
		cost := costs[i]
		if cost == -1 {
			return -1
		} else {
			if maxCost < cost {
				maxCost = cost
			}
		}
	}
	return maxCost
}


//type destCost struct {
//	dest int
//	cost int
//}
//
//type destCosts []destCost
//
//// 邻接表
//func networkDelayTime(times [][]int, n int, k int) int {
//	// prepare
//	costs := make([]int, n+1)
//	for i := range costs {
//		costs[i] = -1
//	}
//	destCostMap := make(map[int]destCosts, len(times))
//	for _, time := range times {
//		s, d, c := time[0], time[1], time[2]
//		if _, ok := destCostMap[s]; !ok {
//			destCostMap[s] = make(destCosts, 0)
//		}
//		destCostMap[s] = append(destCostMap[s], destCost{
//			dest: d,
//			cost: c,
//		})
//	}
//	l := list.New()
//	l.PushBack(k)
//	costs[k] = 0
//	for l.Len() > 0 {
//		curr := l.Remove(l.Front()).(int)
//		dcs, ok := destCostMap[curr]
//		if !ok {
//			continue
//		}
//		for _, dc := range dcs {
//			nc := costs[curr] + dc.cost
//			if costs[dc.dest] == -1 || nc < costs[dc.dest] {
//				costs[dc.dest] = nc
//				l.PushBack(dc.dest)
//			}
//		}
//	}
//	maxCost := 0
//	for i := 1; i < len(costs); i++ {
//		cost := costs[i]
//		if cost == -1 {
//			return -1
//		} else {
//			if maxCost < cost {
//				maxCost = cost
//			}
//		}
//	}
//	return maxCost
//}
