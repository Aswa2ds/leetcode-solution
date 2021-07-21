package offer

func verifyPostorder(postorder []int) bool {
	if len(postorder) == 1 || len(postorder) == 0 {
		return true
	}
	rootVal := postorder[len(postorder)-1]
	i := 0
	for ; i < len(postorder)-1; i++ {
		if postorder[i] > rootVal {
			break
		}
	}
	leftTree := postorder[:i]
	rightTree := postorder[i:len(postorder)-1]
	for _, j := range rightTree {
		if j < rootVal {
			return false
		}
	}
	return verifyPostorder(leftTree) && verifyPostorder(rightTree)
}
