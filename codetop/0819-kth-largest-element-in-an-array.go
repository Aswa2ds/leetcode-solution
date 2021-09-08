package codetop

import (
	"math/rand"
	"time"
)

type bigHeap []int

func (b bigHeap) Len() int {
	return len(b)
}

func (b bigHeap) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b *bigHeap) Swap(i, j int) {
	(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
}

func (b *bigHeap) Push(x interface{}) {
	*b = append(*b, x.(int))
}

func (b *bigHeap) Pop() interface{} {
	ret := (*b)[len(*b)-1]
	*b = (*b)[0 : len(*b)-1]
	return ret
}

//func findKthLargest(nums []int, k int) int {
//	hp := &bigHeap{}
//	heap.Init(hp)
//	size := 0
//	for _, num := range nums {
//		if size < k {
//			heap.Push(hp, num)
//		} else {
//			if num > (*hp)[0] {
//				heap.Push(hp, num)
//				heap.Pop(hp)
//			}
//		}
//	}
//	return (*hp)[0]
//}

func findKthLargest(nums []int, k int) int {
	return divide(nums, 0, len(nums)-1, len(nums)-k)
}

func divide(nums []int, left int, right int, k int) int {
	if left >= right {
		return nums[left]
	}
	key := rand.Intn(right-left) +left
	nums[left], nums[key] = nums[key], nums[left]
	i, j, pivot := left, right, nums[left]
	for i < j {
		for i < j && nums[j] >= pivot {
			j--
		}
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	if i == k {
		return nums[i]
	} else if i < k {
		return divide(nums, i+1, right, k)
	} else {
		return divide(nums, left, i-1, k)
	}

}

func quickSort(nums []int, left, right int) {
	// 终止条件
	if left >= right {
		return
	}
	// 随机key 避免原数组按序排列，导致的n^2 的时间复杂度
	rand.Seed(time.Now().UnixNano())
	key := rand.Intn(right-left)+left
	nums[left], nums[key] = nums[key], nums[left]
	lo, hi, pivot := left, right, nums[left]
	for lo < hi {
		for lo < hi && nums[hi] >= pivot {
			hi--
		}
		for lo < hi && nums[lo] <= pivot {
			lo++
		}
		// 此时lo hi 的位置，要么重合，要么满足交换条件，所以交换就完了
		nums[lo], nums[hi] = nums[hi], nums[lo]
	}
	// 此时lo hi 的位置重合，且不存在越界的情况，且由于lo起始指在pivot 上，不会出现lo 的位置小于pivot的情况，如果后门的数字全都小于pivot，那么此时lo = pivot
	nums[left], nums[lo] = nums[lo], nums[left]
	// 分治左右两侧
	quickSort(nums, left, lo-1)
	quickSort(nums, lo+1, right)
}



//func divide(nums []int, i int, j int, k int) int {
//	pivot := i
//	lo, hi := i+1, j-1
//	for lo < hi {
//		for hi > lo && nums[hi] <= nums[pivot] {
//			hi--
//		}
//		for hi > lo && nums[lo] >= nums[pivot] {
//			lo++
//		}
//		if lo == hi {
//			break
//		}
//		nums[lo], nums[hi] = nums[hi], nums[lo]
//	}
//	if nums[hi] > nums[pivot] {
//		nums[pivot], nums[hi] = nums[hi], nums[pivot]
//		pivot = hi
//	}
//	if pivot == k {
//		return nums[pivot]
//	} else if pivot < k {
//		return divide(nums, pivot+1, j, k)
//	} else {
//		return divide(nums, i, pivot, k)
//	}
//
//}
