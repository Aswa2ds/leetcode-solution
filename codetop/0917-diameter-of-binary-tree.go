package codetop

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var findMaxDiameter func(root *TreeNode) int

	findMaxDiameter = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := findMaxDiameter(root.Left)
		right := findMaxDiameter(root.Right)
		maxDiameter = max(left+right+1, maxDiameter)
		return max(left, right) + 1
	}
	findMaxDiameter(root)
	return maxDiameter - 1
}
