package codetop

func deleteDuplicates(head *ListNode) *ListNode {
	emptyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	tail := emptyHead
	for p := tail.Next; p != nil && p.Next != nil; p = tail.Next {
		q := p.Next
		if p.Val != q.Val {
			tail = tail.Next
		} else {
			for ; q != nil && q.Val == p.Val; q = q.Next {
			}
			tail.Next = q
		}
	}
	return emptyHead.Next
}
