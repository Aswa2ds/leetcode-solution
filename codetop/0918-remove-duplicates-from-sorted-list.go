package codetop

func deleteDuplicates(head *ListNode) *ListNode {
	for p := head; p != nil; {
		q := p.Next
		if q == nil {
			break
		}
		if p.Val == q.Val {
			p.Next = q.Next
		} else {
			p = p.Next
		}
	}
	return head
}
