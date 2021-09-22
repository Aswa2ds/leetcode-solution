package codetop

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	next := head.Next
	next.Next = swapPairs(next.Next)
	head.Next = next.Next
	next.Next = head
	head = next
	return head
}
