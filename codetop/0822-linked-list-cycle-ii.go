package codetop

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast == nil {
		return nil
	}
	s1, s2 := head, slow
	for s1 != s2 {
		s1, s2 = s1.Next, s2.Next
	}
	return s1
}
