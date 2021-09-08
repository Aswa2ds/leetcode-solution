package codetop

func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	list := make([]*TreeNode, 0)
	list = append(list, root)
	for len(list) > 0 {
		line := make([]int, 0)
		n := len(list)
		for i := 0; i < n; i++ {
			node := list[0]
			list = list[1:]
			line = append(line, node.Val)
			if node.Left != nil {
				list = append(list, node.Left)
			}
			if node.Right != nil {
				list = append(list, node.Right)
			}
		}
		ret = append(ret, line)
	}
	return ret
}
