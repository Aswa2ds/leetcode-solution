package codetop

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	c := 0
	emptyHead := &ListNode{}
	tail := emptyHead
	for l1 != nil || l2 != nil || c != 0 {
		node := ListNode{
			Val:  c,
			Next: nil,
		}
		if l1 != nil {
			node.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			node.Val += l2.Val
			l2 = l2.Next
		}
		c = node.Val / 10
		node.Val %= 10
		tail.Next = &node
		tail = tail.Next
	}
	return emptyHead.Next
}
