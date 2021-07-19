package offer

type CQueue struct {
	in, out *stack
}


func Constructor() CQueue {
	return CQueue{
		in:  &stack{},
		out: &stack{},
	}
}


func (this *CQueue) AppendTail(value int)  {
	this.in.push(value)
}


func (this *CQueue) DeleteHead() int {
	if this.out.size() == 0 {
		for this.in.size() > 0 {
			this.out.push(this.in.pop())
		}
	}
	if this.out.size() > 0 {
		return this.out.pop().(int)
	}
	return -1
}

//import "container/list"
//
//type CQueue struct {
//	stack1, stack2 *list.List
//}
//
//
//func Constructor() CQueue {
//	return CQueue{
//		stack1: list.New(),
//		stack2: list.New(),
//	}
//}
//
//
//func (this *CQueue) AppendTail(value int)  {
//	this.stack1.PushBack(value)
//}
//
//
//func (this *CQueue) DeleteHead() int {
//	if this.stack2.Len() == 0 {
//		for this.stack1.Len() > 0 {
//			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
//		}
//	}
//	if this.stack2.Len() != 0 {
//		return this.stack2.Remove(this.stack2.Back()).(int)
//	}
//	return -1
//}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
