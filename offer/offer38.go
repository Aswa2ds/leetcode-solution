package offer

func permutation(s string) []string {
	tmp := make(map[string]bool, 0)
	for i := 0; i < len(s); i++ {
		curr := s[i:i+1]
		strings := permutation(s[:i] + s[i+1:])
		if len(strings) == 0 {
			tmp[curr] = true
		}
		for _, str := range strings {
			tmp[curr+str] = true
		}
	}
	ret := make([]string, 0)
	for key, _ := range tmp {
		ret = append(ret, key)
	}
	return ret
}
