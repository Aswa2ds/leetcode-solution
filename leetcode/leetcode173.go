package leetcode

type BSTIterator struct {
	index int
	list []int
}


func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		index: 0,
		list:  append([]int{-1}, preOrder(root)...),
	}
}

func preOrder(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}
	ret := preOrder(root.Left)
	ret = append(ret, root.Val)
	ret = append(ret, preOrder(root.Right)...)
	return ret
}


func (this *BSTIterator) Next() int {
	this.index++
	return this.list[this.index]
}


func (this *BSTIterator) HasNext() bool {
	return this.index < len(this.list)
}
