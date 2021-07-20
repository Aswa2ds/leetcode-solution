package offer


func deleteNode(head *ListNode, val int) *ListNode {
	emptyHead := new(ListNode)
	emptyHead.Next = head
	s1, s2 := emptyHead, head
	for ; s2 != nil; s1, s2 = s1.Next, s2.Next {
		if s2.Val == val {
			s1.Next = s2.Next
		}
	}
	return emptyHead.Next
}
