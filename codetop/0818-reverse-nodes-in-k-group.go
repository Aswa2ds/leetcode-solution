package codetop

func reverseKGroup(head *ListNode, k int) *ListNode {
	emptyHead := &ListNode{Next: head}
	prev := emptyHead
	for head != nil {
		tail := prev
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return emptyHead.Next
			}
		}
		next := tail.Next
		subReverseList(head, tail)
		prev.Next = tail
		head.Next = next
		prev = head
		head = prev.Next
	}
	return emptyHead.Next
}

func subReverseList(head, tail *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		next := p.Next
		p.Next = prev
		prev = p
		p = next
	}
}