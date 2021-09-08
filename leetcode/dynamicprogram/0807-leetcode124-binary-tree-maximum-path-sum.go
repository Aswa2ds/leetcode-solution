package dynamicprogram

import "math"

func maxPathSum(root *TreeNode) int {
	// 记录的思想体现了dp
	res := math.MinInt32
	var maxPathSumDFS func(*TreeNode) int
	var max func(int, int) int

	max = func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	maxPathSumDFS = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftMaxPathSum := max(maxPathSumDFS(node.Left), 0)
		rightMaxPathSum := max(maxPathSumDFS(node.Right), 0)

		currentNodePathSum := leftMaxPathSum + rightMaxPathSum + node.Val
		res = max(currentNodePathSum, res)
		return max(leftMaxPathSum, rightMaxPathSum) + node.Val
	}

	maxPathSumDFS(root)
	return res
}
