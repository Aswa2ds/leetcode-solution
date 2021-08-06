package stackquiz

type stack []interface{}

func (s *stack) empty() bool {
	return len(*s) == 0
}

func (s *stack) push(v interface{}) {
	*s = append(*s, v)
}

func (s *stack) size() int {
	return len(*s)
}

func (s *stack) pop() interface{} {
	res := (*s)[s.size()-1]
	*s = (*s)[:s.size()-1]
	return res
}

func (s *stack) top() interface{} {
	return (*s)[s.size()-1]
}

func trap(height []int) int {
	n, res := len(height), 0
	s := stack{}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	for i := 0; i < n; i++ {
		for !s.empty() && height[i] > height[s.top().(int)] {
			mid := s.pop().(int)
			if s.empty() {
				break
			}
			res += (min(height[i], height[s.top().(int)]) - height[mid]) * (i - s.top().(int) - 1)
		}
		s.push(i)
	}
	return res
}

// use stack
//func trap(height []int) int {
//	n, res := len(height), 0
//	l := list.New()
//
//	min := func(a, b int) int {
//		if a < b {
//			return a
//		}
//		return b
//	}
//
//	for i := 0; i < n; i++ {
//		for l.Len() != 0 && height[i] > height[l.Back().Value.(int)] {
//			mid := l.Back().Value.(int)
//			l.Remove(l.Back())
//			if l.Len() == 0 {
//				break
//			}
//			res += (min(height[l.Back().Value.(int)], height[i]) - height[mid]) * (i - l.Back().Value.(int) - 1)
//		}
//		l.PushBack(i)
//	}
//	return res
//}

//func trap(height []int) int {
//	n := len(height)
//	res := 0
//	l, r, lmax, rmax := 0, n-1, 0, 0
//	for l < r {
//		if height[l] < height[r] {
//			if height[l] > lmax {
//				lmax = height[l]
//			} else {
//				res += lmax - height[l]
//			}
//			l++
//		} else {
//			if height[r] > rmax {
//				rmax = height[r]
//			} else {
//				res += rmax - height[r]
//			}
//			r--
//		}
//	}
//	return res
//}

// 一列一列算的方式，存储lmax 空间复杂度O(n)
//func trap(height []int) int {
//	n := len(height)
//	lmax := make([]int, n+1)
//
//	max := func(a, b int) int {
//		if a > b {
//			return a
//		}
//		return b
//	}
//
//	min := func(a, b int) int {
//		if a > b {
//			return b
//		}
//		return a
//	}
//
//	for i := 1; i < n; i++ {
//		lmax[i+1] = max(height[i], lmax[i])
//	}
//	rmax, res := 0, 0
//	for i := n-1; i >= 0; i-- {
//		rmax = max(rmax, height[i])
//		res += min(lmax[i+1], rmax) - height[i]
//	}
//	return res
//}