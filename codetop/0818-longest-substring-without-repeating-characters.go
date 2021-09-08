package codetop

func lengthOfLongestSubstring(s string) int {
	s1, length, tmp := -1, 0, 0
	cache := make(map[int32]int)
	for i, ch := range s {
		if index, ok := cache[ch]; ok {
			// exists
			if index >= s1 {
				s1 = index
				tmp = i - s1
				cache[ch] = i
				continue
			}
		}
		cache[ch] = i
		tmp++
		if tmp > length {
			length = tmp
		}
	}
	return length
}
