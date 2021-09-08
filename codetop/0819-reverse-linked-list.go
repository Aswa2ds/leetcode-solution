package codetop

//func reverseList(head *ListNode) *ListNode {
//	if head == nil {
//		return head
//	}
//	tail := head
//	for tail.Next != nil {
//		tail = tail.Next
//	}
//	for head != tail {
//		next := head.Next
//		head.Next = tail.Next
//		tail.Next  = head
//		head = next
//	}
//	return head
//}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	tail := head
	next := tail.Next
	for next != nil {
		tail.Next = next.Next
		next.Next = head
		head = next
		next = tail.Next
	}
	return head
}
