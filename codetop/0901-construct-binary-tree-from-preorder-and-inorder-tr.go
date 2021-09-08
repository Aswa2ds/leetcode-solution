package codetop

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	var rootIndex int
	for i, val := range inorder {
		if val == rootVal {
			rootIndex = i
		}
	}
	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(preorder[1:rootIndex+1], inorder[:rootIndex]),
		Right: buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:]),
	}
}
