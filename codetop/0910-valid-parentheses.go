package codetop

func isValid(s string) bool {
	var stack []byte
	parMap := make(map[byte]byte)
	parMap[')'] = '('
	parMap[']'] = '['
	parMap['}'] = '{'
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b == '(' || b == '[' || b == '{' {
			stack = append(stack, b)
		} else {
			if len(stack) != 0 && stack[len(stack)-1] == parMap[b] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
