package offer

func mirrorTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = mirrorTree(root.Right), mirrorTree(root.Left)
	}
	return root
}
