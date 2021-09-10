package codetop

func isCompleteTree(root *TreeNode) bool {
	reachNil := false
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			reachNil = true
			continue
		}
		if reachNil {
			return false
		}
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}
	return true
}
