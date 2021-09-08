package codetop

import "math"

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ret = math.MinInt64
	max := func(a, b int) int {
	    if a > b {
	        return a
	    }
	    return b
	}
	var subMaxPathSum func(root *TreeNode) int
	subMaxPathSum = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := max(subMaxPathSum(root.Left), 0)
		right := max(subMaxPathSum(root.Right), 0)
		ret = max(left + right + root.Val, ret)
		return max(left, right) + root.Val
	}
	subMaxPathSum(root)
	return ret
}