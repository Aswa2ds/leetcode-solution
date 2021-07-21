package offer

type Direction bool

const (
	FROMLEFTTORIGHT Direction = true
	FROMRIGHTTOLEFT Direction = false
)

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	s := append(make([]*TreeNode, 0), root)
	direction := FROMLEFTTORIGHT
	for len(s) != 0 {
		tmp := make([]int, 0)
		nodeCount := len(s)
		for i := nodeCount-1; i >= 0; i-- {
			node := s[i]
			if node == nil {
				continue
			}
			tmp = append(tmp, node.Val)
			switch direction {
			case FROMLEFTTORIGHT:
				s = append(s, node.Left, node.Right)
			case FROMRIGHTTOLEFT:
				s = append(s, node.Right, node.Left)
			}
		}
		s = s[nodeCount:]
		if len(tmp) != 0 {
			res = append(res, tmp)
		}
		direction = !direction
	}
	return res
}
