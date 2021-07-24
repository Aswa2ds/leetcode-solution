package offer

func preOrder(root *TreeNode) []int {
	var res []int
	if root != nil {
		res = append(res, preOrder(root.Left)...)
		res = append(res, root.Val)
		res = append(res, preOrder(root.Right)...)
	} else {
		res = append(res, 0)
	}
	return res
}

func mirrorPreOrder(root *TreeNode) []int {
	var res []int
	if root != nil {
		res = append(res, mirrorPreOrder(root.Right)...)
		res = append(res, root.Val)
		res = append(res, mirrorPreOrder(root.Left)...)
	}  else {
		res = append(res, 0)
	}
	return res
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left, right := preOrder(root.Left), mirrorPreOrder(root.Right)
	if len(left) != len(right) {
		return false
	}
	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}
