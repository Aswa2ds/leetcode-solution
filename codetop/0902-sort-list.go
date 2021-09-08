package codetop

//import "math"
//
//func sortList(head *ListNode) *ListNode {
//	pre := &ListNode{Next: head}
//	emptyHead := &ListNode{Val: math.MinInt64}
//	for pre.Next != nil {
//		p := pre.Next
//		pre.Next = p.Next
//		p.Next = nil
//
//		q := emptyHead
//		for q.Next != nil {
//			if q.Next.Val > p.Val {
//				break
//			}
//			q = q.Next
//		}
//		p.Next = q.Next
//		q.Next = p
//	}
//	return emptyHead.Next
//}

func sortList(head *ListNode) *ListNode {
	return mergeSortLinkedList(head, nil)
}

func mergeSortLinkedList(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	slow, quick := head, head
	for quick != tail {
		slow = slow.Next
		quick = quick.Next
		if quick != tail {
			quick = quick.Next
		}
	}
	return mergeLinkedList(mergeSortLinkedList(head, slow), mergeSortLinkedList(slow, tail))
}

func mergeLinkedList(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeLinkedList(l1.Next, l2)
		return l1
	}
	l2.Next = mergeLinkedList(l1, l2.Next)
	return l2
}