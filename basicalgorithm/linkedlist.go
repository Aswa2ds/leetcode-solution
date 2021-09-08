package basicalgorithm

type ListNode struct {
	Val int
	Next *ListNode
}

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	p, q := headA, headB
	for p != q {
		if p == nil {
			p = headB
		} else {
			p = p.Next
		}
		if q == nil {
			q = headA
		} else {
			q = q.Next
		}
	}
	return p
}

func ReverseList(head *ListNode) *ListNode {
	tail := head
	next := head.Next
	for next != nil {
		tail.Next = next.Next
		next.Next = head
		head = next
		next = tail.Next
	}
	return head
}

func ReversePartList(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	next := head.Next
	t := head
	for head != tail {
		t.Next = next.Next
		next.Next = head
		head = next
		next = t.Next
	}
	return head, t
}

func IsExistLoop(head *ListNode) bool {
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	c := 0
	p, q := l1, l2
	head := &ListNode{}
	tail := head
	for p != nil && q != nil {
		sum := p.Val + q.Val + c
		val := sum%10
		c = sum/10
		node := &ListNode{Val: val}
		tail.Next = node
		tail = node
		p, q = p.Next, q.Next
	}
	if q != nil {
		p = q
	}
	for p != nil {
		sum := p.Val + c
		val := sum%10
		c = sum/10
		node := &ListNode{Val: val}
		tail.Next = node
		tail = node
		p = p.Next
	}
	if c == 1 {
		tail.Next = &ListNode{Val: 1}
	}
	return head.Next
}

func mergeTwoListsWithNoRecursion(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head
	p, q := l1, l2
	for p != nil && q != nil {
		if p.Val < q.Val {
			tail.Next = p
			p = p.Next
			tail = tail.Next
		} else {
			tail.Next = q
			q = q.Next
			tail = tail.Next
		}
	}
	if q != nil {
		p = q
	}
	tail.Next = p
	return head.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l2
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

func removeElements(head *ListNode, val int) *ListNode {
	emptyHead := &ListNode{Next: head}
	for p := emptyHead; p.Next != nil; {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return emptyHead.Next
}