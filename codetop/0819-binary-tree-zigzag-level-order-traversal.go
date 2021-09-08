package codetop

import "container/list"

func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	left, right := list.New(), list.New()
	right.PushBack(root)
	for right.Len() > 0 {
		line := make([]int, 0)
		for right.Len() > 0 {
			element := right.Back()
			right.Remove(element)
			node := element.Value.(*TreeNode)
			if node.Left != nil {
				left.PushBack(node.Left)
			}
			if node.Right != nil {
				left.PushBack(node.Right)
			}
			line = append(line, node.Val)
		}
		if len(line) > 0 {
			ret = append(ret, line)
		}
		line = make([]int, 0)
		for left.Len() > 0 {
			element := left.Back()
			left.Remove(element)
			node := element.Value.(*TreeNode)
			if node.Right != nil {
				right.PushBack(node.Right)
			}
			if node.Left != nil {
				right.PushBack(node.Left)
			}
			line = append(line, node.Val)
		}
		if len(line) > 0 {
			ret = append(ret, line)
		}
	}
	return ret
}
