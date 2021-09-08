package codetop

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil {
		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
