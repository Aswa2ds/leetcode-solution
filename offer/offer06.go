package offer

func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	list := reversePrint(head.Next)
	return append(list, head.Val)
}

// time: O(n) mem: O(n)
//func reversePrint(head *ListNode) []int {
//	var list []int
//	for p := head; p != nil; p = p.Next {
//		list = append(list, p.Val)
//	}
//	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
//		list[i], list[j] = list[j], list[i]
//	}
//	return list
//}

//func reversePrint(head *ListNode) []int {
//	var newHead *ListNode
//	s1 := head
//	for s1 != nil {
//		s2 := s1.Next
//		s1.Next = newHead
//		newHead = s1
//		s1 = s2
//	}
//	var list []int
//	for p := newHead; p != nil; p = p.Next {
//		list = append(list, p.Val)
//	}
//	return list
//}