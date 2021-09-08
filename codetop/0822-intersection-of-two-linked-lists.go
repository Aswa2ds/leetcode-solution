package codetop

//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	p, q := headA, headB
//	var If func(bool, interface{}, interface{}) interface{}
//	If = func(condition bool, trueValue interface{}, falseValue interface{}) interface{} {
//		if condition {
//			return trueValue
//		}
//		return falseValue
//	}
//	for p != q {
//		p = If(p == nil, headB, p.Next).(*ListNode)
//		q = If(q == nil, headA, q.Next).(*ListNode)
//	}
//	return p
//}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
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