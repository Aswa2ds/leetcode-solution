package codetop

type heap []*ListNode

func (hp *heap) down(i int) {
	n := len(*hp)
	for {
		j1 := 2*i + 1
		if j1 >= n {
			break
		}
		j := j1
		j2 := j1 + 1
		if j2 < n && (*hp)[j2].Val < (*hp)[j1].Val {
			j = j2
		}
		if (*hp)[i].Val > (*hp)[j].Val {
			(*hp)[i], (*hp)[j] = (*hp)[j], (*hp)[i]
		} else {
			break
		}
		i = j
	}
}

func (hp *heap) up(j int) {
	for j != 0 {
		i := (j - 1) >> 1
		if (*hp)[i].Val > (*hp)[j].Val {
			(*hp)[i], (*hp)[j] = (*hp)[j], (*hp)[i]
		}
		j = i
	}
}

func (hp *heap) push(v *ListNode) {
	*hp = append(*hp, v)
	hp.up(len(*hp) - 1)
}

func (hp *heap) pop() *ListNode {
	ret := (*hp)[0]
	(*hp)[0] = (*hp)[len((*hp))-1]
	*hp = (*hp)[:len((*hp))-1]
	hp.down(0)
	return ret
}

func mergeKLists(lists []*ListNode) *ListNode {
	emptyHead := &ListNode{}
	tail := emptyHead
	hp := make(heap, 0)
	for i := range lists {
		if lists[i] != nil {
			hp.push(lists[i])
		}
	}
	for len(hp) != 0 {
		node := hp.pop()
		if node.Next != nil {
			hp.push(node.Next)
		}
		tail.Next = node
		tail = tail.Next
		tail.Next = nil
	}
	return emptyHead.Next
}
