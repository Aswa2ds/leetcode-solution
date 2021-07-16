package offer

func replaceSpace(s string) string {
	bytes := []byte(s)
	res := make([]byte, 0)
	for _, b := range bytes {
		if b == ' ' {
			res = append(res, '%')
			res = append(res, '2')
			res = append(res, '0')
		} else {
			res = append(res, b)
		}
	}
	return string(res)
}