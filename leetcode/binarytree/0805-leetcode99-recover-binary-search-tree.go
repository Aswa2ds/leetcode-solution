package binarytree

func recoverTree(root *TreeNode)  {
	var x, y, pred, predecessor *TreeNode
	for root != nil {
		if root.Left != nil {
			predecessor = root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				predecessor.Right = root
				root = root.Left
			} else {
				if pred != nil && pred.Val > root.Val {
					x = root
					if y == nil {
						y = pred
					}
				}
				pred = root
				root = root.Right
				predecessor.Right = nil
			}
		} else {
			if pred != nil && pred.Val > root.Val {
				x = root
				if y == nil {
					y = pred
				}
			}
			pred = root
			root = root.Right
		}
	}
	x.Val, y.Val = y.Val, x.Val
}

//func recoverTree(root *TreeNode)  {
//	stack := []*TreeNode
//	var x, y, prev *TreeNode
//	for {
//		for root != nil {
//			stack = append(stack, root)
//			root = root.Left
//		}
//		root = stack[len(stack)-1]
//		stack = stack[:len(stack)-1]
//		if prev != nil && root.Val < prev.Val {
//			x = root
//			if y != nil {
//				y = prev
//			} else {
//				break
//			}
//		}
//		prev = root
//		root = root.Right
//	}
//	x.Val, y.Val = y.Val, x.Val
//}
