package codetop

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	emptyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	slow, fast := emptyHead, emptyHead
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return emptyHead.Next
}
