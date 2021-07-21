package offer

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	m[nil] = nil
	emptyHead := new(Node)
	tail := emptyHead
	for p := head; p != nil; p = p.Next {
		newNode := new(Node)
		newNode.Val = p.Val
		tail.Next = newNode
		tail = tail.Next
		m[p] = newNode
	}
	newHead := emptyHead.Next
	for p, q := head, newHead; p != nil; p, q = p.Next, q.Next {
		q.Random = m[p.Random]
	}

	return newHead
}
