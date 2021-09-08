package codetop

import "math"

//func isValidBST(root *TreeNode) bool {
//	return InOrder(new([]int), root)
//}
//
//func InOrder(res *[]int, root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//	if !InOrder(res, root.Left) {
//		return false
//	}
//	if len(*res) > 0 && (*res)[len(*res)-1] >= root.Val {
//		return false
//	}
//	*res = append(*res, root.Val)
//	if !InOrder(res, root.Right) {
//		return false
//	}
//	return true
//}

func isValidBST(root *TreeNode) bool {
	pre := math.MinInt64
	return InOrder(&pre, root)
}

func InOrder(res *int, root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !InOrder(res, root.Left) {
		return false
	}
	if *res >= root.Val {
		return false
	}
	*res = root.Val
	if !InOrder(res, root.Right) {
		return false
	}
	return true
}
