package codetop

func flatten(root *TreeNode) {
	for node := root; node != nil; node = node.Right {
		if node.Left == nil {
			continue
		}
		prev := node.Left
		for prev.Right != nil {
			prev = prev.Right
		}
		prev.Right = node.Right
		node.Right = node.Left
		node.Left = nil
	}
}
