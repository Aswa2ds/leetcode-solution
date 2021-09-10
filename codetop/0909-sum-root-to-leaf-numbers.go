package codetop

func sumNumbers(root *TreeNode) int {
	ret := 0
	var calSum func(root *TreeNode, base int)
	isLeave := func(root *TreeNode) bool {
		if root.Left == nil && root.Right == nil {
			return true
		}
		return false
	}
	calSum = func(root *TreeNode, base int) {
		if root == nil {
			return
		}
		base = base + root.Val
		if isLeave(root) {
			ret += base
		}
		calSum(root.Left, base*10)
		calSum(root.Right, base*10)
	}
	calSum(root, 0)
	return ret
}
