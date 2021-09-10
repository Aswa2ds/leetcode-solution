package codetop

func pathSum(root *TreeNode, targetSum int) [][]int {
	ret := make([][]int, 0)
	var dfs func(root *TreeNode, target int, base []int)
	dfs = func(root *TreeNode, target int, base []int) {
		if root == nil {
			return
		}
		target = target - root.Val
		base = append(base, root.Val)
		if root.Left == nil && root.Right == nil {
			if target == 0 {
				cp := make([]int, len(base))
				copy(cp, base)
				ret = append(ret, cp)
			}
			return
		}
		dfs(root.Left, target, base)
		dfs(root.Right, target, base)
	}
	dfs(root, targetSum, []int{})
	return ret
}
