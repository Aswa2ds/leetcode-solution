package offer

func levelOrder_I(root *TreeNode) []int {
	var queue []*TreeNode
	var ret []int
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			continue
		}
		ret = append(ret, node.Val)
		queue = append(queue, node.Left, node.Right)
	}
	return ret
}
