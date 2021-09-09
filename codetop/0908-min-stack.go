package codetop

type MinStack struct {
	s    []int
	mins []int
}

// func Constructor() MinStack {
// 	return MinStack{
// 		s:    []int{},
// 		mins: []int{},
// 	}
// }

func (this *MinStack) Push(val int) {
	this.s = append(this.s, val)
	if len(this.mins) == 0 || this.mins[len(this.mins)-1] >= val {
		this.mins = append(this.mins, val)
	}
}

func (this *MinStack) Pop() {
	val := this.s[len(this.s)-1]
	this.s = this.s[:len(this.s)-1]
	if val == this.mins[len(this.mins)-1] {
		this.mins = this.mins[:len(this.mins)-1]
	}
}

func (this *MinStack) Top() int {
	return this.s[len(this.s)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}
