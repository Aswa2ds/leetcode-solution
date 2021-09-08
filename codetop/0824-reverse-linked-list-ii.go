package codetop

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	emptyHead := &ListNode{Next: head}
	pre := emptyHead
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := left; i < right; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = cur
		pre.Next = next
	}
	return emptyHead.Next
}
