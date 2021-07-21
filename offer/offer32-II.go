package offer

func levelOrder_II(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	queue := append(make([]*TreeNode, 0), root)
	for len(queue) != 0 {
		tmp := make([]int, 0)
		nodeCount := len(queue)
		for i := 0; i < nodeCount; i++ {
			node := queue[i]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[nodeCount:]
		res = append(res, tmp)
	}
	return res
}

//func levelOrder(root *TreeNode) [][]int {
//	if root == nil {
//		return [][]int{}
//	}
//	res := make([][]int, 0)
//	depth := 0
//	queue1 := make([]*TreeNode, 0)
//	queue2 := make([]*TreeNode, 0)
//	queue1 = append(queue1, root)
//	for len(queue1) != 0 || len(queue2) != 0 {
//		if len(queue1) != 0 {
//			res = append(res, make([]int, 0))
//			for len(queue1) != 0 {
//				node := queue1[0]
//				queue1 = queue1[1:]
//				res[depth] = append(res[depth], node.Val)
//				if node.Left != nil {
//					queue2 = append(queue2, node.Left)
//				}
//				if node.Right != nil {
//					queue2 = append(queue2, node.Right)
//				}
//			}
//			depth++
//		}
//
//		if len(queue2) != 0 {
//			res = append(res, make([]int, 0))
//			for len(queue2) != 0 {
//				node := queue2[0]
//				queue2 = queue2[1:]
//				res[depth] = append(res[depth], node.Val)
//				if node.Left != nil {
//					queue1 = append(queue1, node.Left)
//				}
//				if node.Right != nil {
//					queue1 = append(queue1, node.Right)
//				}
//			}
//			depth++
//		}
//	}
//	return res
//}
