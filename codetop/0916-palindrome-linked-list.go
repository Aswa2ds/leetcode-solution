package codetop

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// find second half start
	slow, fast := head, head
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		slow = slow.Next
	}
	secondHalfStart := slow
	var prev, cur *ListNode = nil, secondHalfStart
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	secondHalfStart = prev
	ret := true
	for p1, p2 := head, secondHalfStart; ret && p2 != nil; p1, p2 = p1.Next, p2.Next {
		if p1.Val != p2.Val {
			ret = false
		}
	}
	return ret
}
