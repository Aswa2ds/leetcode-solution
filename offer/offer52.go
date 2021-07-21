package offer

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p, q := headA, headB
	for p != q {
		if p == nil {
			p = headB
		} else {
			p = p.Next
		}
		if q == nil {
			q = headA
		} else {
			q = q.Next
		}
	}
	return p
}

//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	var lenA, lenB int
//	for p := headA; p != nil; p, lenA = p.Next, lenA+1 { }
//	for q := headB; q != nil; q, lenB = q.Next, lenB+1 { }
//	if lenB < lenA {
//		lenA, lenB = lenB, lenA
//		headA, headB = headB, headA
//	}
//	for lenA < lenB {
//		newHead := new(ListNode)
//		newHead.Next = headA
//		headA = newHead
//		lenA++
//	}
//	p, q := headA, headB
//	for ; p != q && p != nil; p, q = p.Next, q.Next {
//	}
//	return p
//}
