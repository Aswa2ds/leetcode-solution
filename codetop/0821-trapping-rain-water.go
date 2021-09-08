package codetop

//type stack []int
//
//func (s *stack) isEmpty() bool {
//	return s.len() == 0
//}
//
//func (s *stack) push(v int) {
//	*s = append(*s, v)
//}
//
//func (s *stack) pop() int {
//	ret := (*s)[s.len()-1]
//	*s = (*s)[:s.len()-1]
//	return ret
//}
//
//func (s *stack) top() int {
//	return (*s)[s.len()-1]
//}
//
//func (s *stack) len() int {
//	return len(*s)
//}
//
//func trap(height []int) int {
//	s := make(stack, 0)
//	res := 0
//	min := func(a, b int) int {
//	    if a < b {
//	    	return a
//		}
//		return b
//	}
//	for i, h := range height {
//		for !s.isEmpty() && h > height[s.top()] {
//			mid := height[s.pop()]
//			if s.isEmpty() {
//				break
//			}
//			res += (min(height[s.top()], h)-mid) * (i-s.top()-1)
//		}
//		s.push(i)
//	}
//	return res
//}


func trap(height []int) int {
	lmax := make([]int, len(height)+1)

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 0; i < len(height); i++ {
		lmax[i+1] = max(lmax[i], height[i])
	}
	rmax, res := 0, 0
	for i := len(height)-1; i >= 0; i-- {
		rmax = max(rmax, height[i])
		res += min(rmax, lmax[i+1]) - height[i]
	}
	return res
}
