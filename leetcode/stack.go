package offer

type stack []interface{}

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
