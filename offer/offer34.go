package offer

func subPathSum(root *TreeNode, target int, path []int, res *[][]int) {
	if root == nil {
		return
	}
	if target == root.Val {
		if root.Left == nil && root.Right == nil {
			*res = append(*res, append(make([]int, 0), path...))
			return
		}
	}
	target -= root.Val
	subPathSum(root.Left, target, append(path, root.Val), res)
	subPathSum(root.Right, target, append(path, root.Val), res)
}

func pathSum(root *TreeNode, target int) [][]int {
	res := new([][]int)
	subPathSum(root, target, make([]int, 0), res)
	return *res
}


//func pathSum(root *TreeNode, target int) [][]int {
//	res := make([][]int, 0)
//	path := []int{}
//	var dfs func(*TreeNode, int)
//	dfs = func(node *TreeNode, left int) {
//		if node == nil{
//			return
//		}
//		left -= node.Val
//		path = append(path, node.Val)
//		defer func() {
//			path = path[:len(path)-1]
//		}()
//		if left == 0 && node.Left == nil && node.Right == nil {
//			res = append(res, append(make([]int, 0), path...))
//			return
//		}
//		dfs(node.Left, left)
//		dfs(node.Right, left)
//	}
//	dfs(root, target)
//	return res
//}
