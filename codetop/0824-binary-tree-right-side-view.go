package codetop

func rightSideView(root *TreeNode) []int {
	var dfs func(*TreeNode, int, int) ([]int, int)
	dfs = func(root *TreeNode, currDepth, rightMaxDepth int) ([]int, int) {
		if root == nil {
			return make([]int, 0), rightMaxDepth
		}
		ret := make([]int, 0)
		if currDepth > rightMaxDepth {
			ret = append(ret, root.Val)
			rightMaxDepth = currDepth
		}
		right, rightMaxDepth := dfs(root.Right, currDepth+1, rightMaxDepth)
		left, rightMaxDepth := dfs(root.Left, currDepth+1, rightMaxDepth)
		ret = append(ret, right...)
		ret = append(ret, left...)
		return ret, rightMaxDepth
	}
	ret, _ := dfs(root, 1, 0)
	return ret
}
