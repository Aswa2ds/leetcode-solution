package offer

func validateStackSequences(pushed []int, popped []int) bool {
	tmp := make([]int, 0)
	for _, pop := range popped {
		if len(tmp) != 0 && tmp[len(tmp)-1] == pop {
			tmp = tmp[:len(tmp)-1]
		} else {
			flag := false
			for j, push := range pushed {
				if push == pop {
					tmp = append(tmp, pushed[:j]...)
					pushed = pushed[j+1:]
					flag = true
					break
				}
			}
			if !flag {
				return false
			}
		}
	}
	return true
}
