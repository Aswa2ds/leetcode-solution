package offer

func getKthFromEnd(head *ListNode, k int) *ListNode {
	p := head
	for ; k > 0; p, k = p.Next, k-1 {}
	q := head
	for p != nil {
		p, q = p.Next, q.Next
	}
	return q
}
