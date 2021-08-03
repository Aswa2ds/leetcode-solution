package leetcode

import (
	"fmt"
	"math"
)

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	ret := make([]int, 0)
	distanceKBelow(target, k, &ret)
	distanceKAbove(root, target, k, &ret)
	return ret
}

func distanceKAbove(root *TreeNode, target *TreeNode, k int, ret *[]int) int {
	if k == -1 {
		return -1
	}
	if root == nil {
		return math.MinInt32
	}
	if root == target {
		return k-1
	}
	left := distanceKAbove(root.Left, target, k, ret)
	right := distanceKAbove(root.Right, target, k, ret)
	fmt.Println(left)
	fmt.Println(right)
	if left > 0 {
		distanceKBelow(root.Right, left-1, ret)
		return left-1
	}
	if right > 0 {
		distanceKBelow(root.Left, right-1, ret)
		return right-1
	}
	if left == 0 || right == 0 {
		distanceKBelow(root, 0, ret)
		return -1
	}
	return -1
}

func distanceKBelow(root *TreeNode, k int, ret *[]int) {
	if root == nil {
		return
	}
	if k == 0 {
		*ret = append(*ret, root.Val)
		return
	}
	distanceKBelow(root.Left, k-1, ret)
	distanceKBelow(root.Right, k-1, ret)

}


