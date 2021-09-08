package codetop

//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	emptyHead := &ListNode{}
//	tail := emptyHead
//	p, q := l1, l2
//	for p != nil && q != nil {
//		if p.Val < q.Val {
//			tail.Next = p
//			p = p.Next
//		} else {
//			tail.Next = q
//			q = q.Next
//		}
//		tail = tail.Next
//	}
//	if q != nil {
//		p = q
//	}
//	tail.Next = p
//	return emptyHead.Next
//}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}