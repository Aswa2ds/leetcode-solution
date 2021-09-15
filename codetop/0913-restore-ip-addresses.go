package codetop

import "strconv"

func restoreIpAddresses(s string) []string {
	ret := []string{}
	segments := [4]int{}
	var dfs func(int, int)
	dfs = func(segId, segStart int) {
		if segId == 4 {
			if segStart == len(s) {
				ipAddr := ""
				for _, segment := range segments {
					ipAddr += strconv.Itoa(segment) + "."
				}
				ret = append(ret, ipAddr[:len(ipAddr)-1])
			}
			return
		}
		if segStart == len(s) {
			return
		}

		if s[segStart] == '0' {
			segments[segId] = 0
			dfs(segId+1, segStart+1)
			return
		}
		addr := 0
		for segEnd := segStart; segEnd < len(s); segEnd++ {
			addr = addr*10 + int(s[segEnd]-'0')
			if addr <= 0xFF {
				segments[segId] = addr
				dfs(segId+1, segEnd+1)
			} else {
				break
			}
		}
	}
	dfs(0, 0)
	return ret
}
