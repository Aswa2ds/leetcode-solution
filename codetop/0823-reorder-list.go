package codetop

import "fmt"

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	l1, l2 := head, slow.Next
	slow.Next = nil

	reverseList := func(head *ListNode) *ListNode {
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
	l2 = reverseList(l2)
	for p := l2; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}

	emptyHead := &ListNode{}
	tail := emptyHead
	for l1 != nil && l2 != nil {
		tail.Next = l1
		tail = tail.Next
		l1 = tail.Next
		tail.Next = l2
		tail = tail.Next
		l2 = l2.Next
	}
	tail.Next = l1

}
