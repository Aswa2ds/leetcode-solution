package offer

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val:   preorder[0],
	}
	var i int
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:i+1], inorder[0:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}

//func buildTree(preorder []int, inorder []int) *TreeNode {
//	//if len(preorder) == 0 {
//	//	return nil
//	//}
//	if len(preorder) == 1 {
//		return &TreeNode{
//			Val:   preorder[0],
//			Left:  nil,
//			Right: nil,
//		}
//	}
//	rootVal := preorder[0]
//	var leftSubTreeNodeCount int
//	for i, val := range inorder {
//		if val == rootVal {
//			leftSubTreeNodeCount = i
//			break
//		}
//	}
//	leftPreorder := preorder[1:1+leftSubTreeNodeCount]
//	leftInorder := inorder[0:leftSubTreeNodeCount]
//	rightPreorder := preorder[leftSubTreeNodeCount+1:]
//	rightInorder := inorder[leftSubTreeNodeCount+1:]
//	return &TreeNode{
//		Val: rootVal,
//		Left: buildTree(leftPreorder, leftInorder),
//		Right: buildTree(rightPreorder, rightInorder),
//	}
//}
