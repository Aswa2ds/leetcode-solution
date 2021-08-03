package offer

import "math"

type MedianFinder struct {
	length int
	list   []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		length: 0,
		list:   []int{math.MinInt32, math.MaxInt32},
	}
}

func (this *MedianFinder) AddNum(num int) {
	low, high := 0, this.length+1
	for low < high {
		mid := (low + high) / 2
		if this.list[mid] < num {
			low = mid + 1
		} else if this.list[mid] > num {
			high = mid
		} else {
			low, high = mid, mid
		}
	}
	newList := make([]int, this.length+3)
	for i := range this.list {
		if i < low {
			newList[i] = this.list[i]
		} else if i == low {
			newList[i] = num
		} else {
			newList[i] = this.list[i-1]
		}
	}
	this.list = newList
	this.length++
}

func (this *MedianFinder) FindMedian() float64 {
	if this.length == 0 {
		return 0
	}
	if (this.length & 1) == 1 {
		midIndex := (this.length+1) >> 1
		return float64(this.list[midIndex])
	} else {
		midIndex1 := (this.length) >> 1
		return (float64(this.list[midIndex1]) + float64(this.list[midIndex1+1]))/2
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
