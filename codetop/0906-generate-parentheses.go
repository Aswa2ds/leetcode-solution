package codetop

func generateParenthesis(n int) []string {
	var left, right int
	ret := make([]string, 0)
	genPar(&ret, "", left, right, n)
	return ret
}

func genPar(ret *[]string, base string, left, right, n int) {
	if left == n {
		for right < n {
			base += ")"
			right++
		}
		*ret = append(*ret, base)
		return
	}
	if left == right {
		genPar(ret, base+"(", left+1, right, n)
	} else {
		genPar(ret, base+"(", left+1, right, n)
		genPar(ret, base+")", left, right+1, n)
	}
}
