package offer

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var newHead, s1, s2 *ListNode = nil, head, head.Next
	for s2 != nil {
		s1.Next = newHead
		newHead = s1
		s1, s2 = s2, s2.Next
	}
	s1.Next = newHead
	newHead = s1
	return newHead
}
