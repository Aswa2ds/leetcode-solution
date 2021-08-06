package everyday

func eventualSafeNodes(graph [][]int) []int {
	safety := make([]int, len(graph))
	for i, line := range graph {
		if len(line) == 0 {
			safety[i] = 2
		}
	}
	if len(safety) == 0 {
		return make([]int, 0)
	}
	res := make([]int, 0)
	for i := range graph {
		if ensureSafety(graph, i, safety) {
			res = append(res, i)
		}
	}
	return res
}

// 我的方法是，将已经遍历的数据存储到stack 里，不仅增加的空间复杂度，且在调用contains 时，需要遍历stack
// 我开始提交的是在判断好点状态后不break，后续的contains 调用的时间会更高
//func ensureSafety(graph [][]int, src int, s stack, safety []int) {
//	if s.contains(src) {
//		safety[src] = 1
//	}
//	if safety[src] == 1 {
//		return
//	}
//	if safety[src] == 2 {
//		return
//	}
//	s.push(src)
//	defer s.pop()
//	for _, desc := range graph[src] {
//		ensureSafety(graph, desc, s, safety)
//		if safety[desc] == 1 {
//			safety[src] = 1
//		}
//	}
//	if safety[src] == 0 {
//		safety[src] = 2
//	}
//}
// 后来将break 添加，由于本身有剪枝，反而降低了contains 的调用时间，时间复杂度没变，但是稍有优化
//func ensureSafety(graph [][]int, src int, s stack, safety []int) {
//	if s.contains(src) {
//		safety[src] = 1
//	}
//	if safety[src] == 1 {
//		return
//	}
//	if safety[src] == 2 {
//		return
//	}
//	s.push(src)
//	defer s.pop()
//	for _, desc := range graph[src] {
//		ensureSafety(graph, desc, s, safety)
//		if safety[desc] == 1 {
//			safety[src] = 1
//            break
//		}
//	}
//	if safety[src] == 0 {
//		safety[src] = 2
//	}
//}

func ensureSafety(graph [][]int, src int, safety []int) bool {
	if safety[src] == 1 {
		return false
	}
	if safety[src] == 2 {
		return true
	}
	safety[src] = 1
	for _, desc := range graph[src] {
		if !ensureSafety(graph, desc, safety) {
			return false
		}
	}
	safety[src] = 2
	return true
}
