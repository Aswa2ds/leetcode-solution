package codetop

type stack []int

func (s *stack) push(val int) {
	*s = append(*s, val)
}

func (s *stack) pop() int {
	ret := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ret
}

func (s *stack) top() int {
	ret := (*s)[len(*s)-1]
	return ret
}

type MyQueue struct {
	s1, s2 stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	queue := MyQueue{s1: stack{}, s2: stack{}}
	return queue
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.s1.push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}
	if len(this.s2) == 0 {
		for len(this.s1) != 0 {
			this.s2.push(this.s1.pop())
		}
	}
	return this.s2.pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}
	if len(this.s2) == 0 {
		for len(this.s1) != 0 {
			this.s2.push(this.s1.pop())
		}
	}
	return this.s2.top()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.s1) == 0 && len(this.s2) == 0
}
