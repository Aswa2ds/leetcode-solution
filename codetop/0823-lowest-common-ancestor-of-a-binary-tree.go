package codetop

//func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//	pLeftAncestor := isAncestor(root.Left, p)
//	pRightAncestor := isAncestor(root.Right, p)
//	qLeftAncestor := isAncestor(root.Left, q)
//	qRightAncestor := isAncestor(root.Right, q)
//	if (pLeftAncestor && qRightAncestor) || (pRightAncestor && qLeftAncestor) {
//		return root
//	} else if pLeftAncestor && qLeftAncestor {
//		return lowestCommonAncestor(root.Left, p, q)
//	} else if pRightAncestor && qRightAncestor {
//		return lowestCommonAncestor(root.Right, p, q)
//	} else {
//		return root
//	}
//}
//
//func isAncestor(root, p *TreeNode) bool {
//	if root == nil {
//		return false
//	}
//	if root == p {
//		return true
//	}
//	return isAncestor(root.Left, p) || isAncestor(root.Right, p)
//}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p == root || q == root {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left

}
