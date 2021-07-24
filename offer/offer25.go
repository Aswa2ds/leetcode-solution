package offer

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := new(ListNode)
	tail := newHead
	s1, s2 := l1, l2
	for s1 != nil && s2 != nil {
		if s1.Val <= s2.Val {
			tail.Next = s1
			tail = tail.Next
			s1 = s1.Next
			tail.Next = nil
		}
		if s1 == nil {
			break
		}
		if s1.Val > s2.Val {
			tail.Next = s2
			tail = tail.Next
			s2 = s2.Next
			tail.Next = nil
		}
		if s2 == nil {
			break
		}
	}
	if s1 == nil {
		s1 = s2
	}
	tail.Next = s1
	return newHead.Next
}
