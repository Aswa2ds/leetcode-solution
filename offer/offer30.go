package offer

type MinStack struct {
	list []int
	minList []int
}

/** initialize your data structure here. */
//func Constructor() MinStack {
//	return MinStack{
//		list:    make([]int, 0),
//		minList: make([]int, 0),
//	}
//}


func (this *MinStack) Push(x int)  {
	this.list = append(this.list, x)
	if len(this.minList) == 0 || x <= this.minList[len(this.minList)-1] {
		this.minList = append(this.minList, x)
	}
}


func (this *MinStack) Pop()  {
	if len(this.list) == 0 {
		return
	}
	x := this.list[len(this.list)-1]
	this.list = this.list[:len(this.list)-1]
	if this.minList[len(this.minList)-1] == x {
		this.minList = this.minList[:len(this.minList)-1]
	}
}


func (this *MinStack) Top() int {
	if len(this.list) == 0 {
		return 0
	}
	return this.list[len(this.list)-1]
}


func (this *MinStack) Min() int {
	if len(this.minList) == 0 {
		return 0
	}
	return this.minList[len(this.minList)-1]
}
