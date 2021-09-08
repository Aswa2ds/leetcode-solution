package basicalgorithm

import (
	"container/list"
	"math"
	"math/rand"
	"time"
)

func BubbleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := len(nums) - 1; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
}

func BubbleSortWithFlag(nums []int) {
	if len(nums) <= 1 {
		return
	}
	flag := true
	for i := len(nums) - 1; i >= 1; i-- {
		if flag {
			flag = false
			for j := 1; j <= i; j++ {
				if nums[j-1] > nums[j] {
					nums[j-1], nums[j] = nums[j], nums[j-1]
					flag = true
				}
			}
		}
	}
}

func BubbleSortWithRecord(nums []int) {
	if len(nums) <= 1 {
		return
	}
	record := len(nums) - 1
	for record > 0 {
		border := record
		record = -1
		for i := 1; i <= border; i++ {
			if nums[i-1] > nums[i] {
				nums[i-1], nums[i] = nums[i], nums[i-1]
				record = i
			}
		}
	}
}

func CockTailSort(nums []int) {
	if len(nums) < 1 {
		return
	}
	for i := 0; i < len(nums)/2; i++ {
		isSorted := true
		j := i
		for ; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				isSorted = false
			}
		}
		if isSorted {
			break
		}
		isSorted = true
		for ; j > i; j-- {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
				isSorted = false
			}
		}
		if isSorted {
			break
		}
	}
}

func SelectSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		minIndex := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}

func InsertSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i; j >= 1; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			} else {
				break
			}
		}
	}
}

func InsertSortWithQuickSearch(nums []int) {
	for i := 0; i < len(nums); i++ {
		copyNums := make([]int, len(nums))
		copy(copyNums, nums)
		searchPart := copyNums[0:i]
		leftPart := copyNums[i+1:]
		target := copyNums[i]
		index := quickSearch(searchPart, target, 0, i)
		nums = nums[:0]
		for j := 0; j < index; j++ {
			nums = append(nums, searchPart[j])
		}
		nums = append(nums, target)
		for j := index; j < len(searchPart); j++ {
			nums = append(nums, searchPart[j])
		}
		nums = append(nums, leftPart...)
	}
}

func quickSearch(nums []int, target int, left int, right int) int {
	if left >= right {
		return left
	}
	mid := (left + right) >> 1
	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		return quickSearch(nums, target, mid+1, right)
	} else {
		return quickSearch(nums, target, left, mid)
	}
}

func QuickSort(nums []int) {
	lo, hi := 0, len(nums)-1
	var qs func(nums []int, lo, hi int)
	qs = func(nums []int, lo, hi int) {
		if lo >= hi {
			return
		}
		//pivotIndex := partitionSingleSide(nums, lo, hi)
		pivotIndex := partitionDoubleSides(nums, lo, hi)
		qs(nums, lo, pivotIndex-1)
		qs(nums, pivotIndex+1, hi)
	}
	qs(nums, lo, hi)
}

func partitionSingleSide(nums []int, left, right int) int {
	rand.Seed(time.Now().UnixNano())
	mark := left
	key := rand.Intn(right-left) + left
	pivot := nums[key]
	nums[key], nums[left] = nums[left], nums[key]
	for i := left + 1; i <= right; i++ {
		if nums[i] < pivot {
			mark++
			nums[mark], nums[i] = nums[i], nums[mark]
		}
	}
	nums[left] = nums[mark]
	nums[mark] = pivot
	return mark
}

func partitionDoubleSides(nums []int, left, right int) int {
	rand.Seed(time.Now().UnixNano())
	key := rand.Intn(right-left) + left
	pivot := nums[key]
	nums[key], nums[left] = nums[left], nums[key]
	lo, hi := left, right
	for lo < hi {
		for lo < hi && nums[hi] >= pivot {
			hi--
		}
		for lo < hi && nums[lo] <= pivot {
			lo++
		}
		nums[lo], nums[hi] = nums[hi], nums[lo]
	}
	nums[left] = nums[lo]
	nums[lo] = pivot
	return lo
}

func QuickSortNoRecursion(nums []int) {
	lo, hi := 0, len(nums)-1
	var qs func(nums []int, lo, hi int)
	qs = func(nums []int, lo, hi int) {
		stack := list.New()
		stack.PushBack(lo)
		stack.PushBack(hi)
		for stack.Len() > 0 {
			leftElement := stack.Front()
			stack.Remove(leftElement)
			left := leftElement.Value.(int)
			rightElement := stack.Front()
			stack.Remove(rightElement)
			right := rightElement.Value.(int)
			if left >= right {
				continue
			}
			pivotIndex := partitionSingleSide(nums, left, right)
			stack.PushBack(left)
			stack.PushBack(pivotIndex - 1)
			stack.PushBack(pivotIndex + 1)
			stack.PushBack(right)
		}
	}
	qs(nums, lo, hi)
}

func HeapSort(nums []int) {
	adjustHeap(nums, len(nums))
	for i := range nums {
		popToEnd(nums, len(nums)-i)
	}
}

func popToEnd(nums []int, size int) int {
	ret := nums[0]
	nums[0], nums[size-1] = nums[size-1], nums[0]
	shiftDown(nums, 0, size-1)
	return ret
}

func adjustHeap(nums []int, size int) {
	n := (size - 1) / 2
	for i := n; i >= 0; i-- {
		shiftDown(nums, i, size)
	}
}

func shiftDown(nums []int, i, size int) {
	j1 := 2*i + 1
	j2 := 2*i + 2
	if j1 > size-1 {
		return
	}
	j := j1
	if j2 < size {
		if nums[j2] > nums[j] {
			j = j2
		}
	}
	if nums[i] < nums[j] {
		nums[i], nums[j] = nums[j], nums[i]
	}
	shiftDown(nums, j, size)
}

func ShellSort(nums []int) {
	for gap := len(nums) >> 1; gap > 0; gap = gap >> 1 {
		for i := gap; i < len(nums); i++ {
			for j := i; j-gap >= 0 && nums[j-gap] > nums[j]; j = j - gap {
				nums[j-gap], nums[j] = nums[j], nums[j-gap]
			}
		}
	}
}

func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	lo, hi := 0, len(nums)-1
	mid := (lo + hi + 1) >> 1
	return merge(MergeSort(nums[:mid]), MergeSort(nums[mid:]))
}

func merge(left []int, right []int) []int {
	ret := make([]int, 0)
	s1, s2 := 0, 0
	for s1 < len(left) && s2 < len(right) {
		if left[s1] < right[s2] {
			ret = append(ret, left[s1])
			s1++
		} else {
			ret = append(ret, right[s2])
			s2++
		}
	}
	for ; s1 < len(left); s1++ {
		ret = append(ret, left[s1])
	}
	for ; s2 < len(right); s2++ {
		ret = append(ret, right[s2])
	}
	return ret
}

func CountSort(nums []int) {
	max := math.MinInt64
	base := math.MaxInt64
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < base {
			base = num
		}
	}
	offset := max - base
	counter := make([]int, offset+1)
	for _, num := range nums {
		counter[num-base]++
	}
	j := 0
	for i, count := range counter {
		for count > 0 {
			nums[j] = i + base
			j++
			count--
		}
	}
}

func BucketSort(nums []int) {
	max := math.MinInt64
	base := math.MaxInt64
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < base {
			base = num
		}
	}
	offset := max - base
	bucketNum := 5
	interval := offset / bucketNum
	buckets := make([][]int, bucketNum+1)
	for i := range buckets {
		buckets[i] = make([]int, 0)
	}
	for _, num := range nums {
		bucketIndex := num / interval
		if bucketIndex >= bucketNum {
			bucketIndex = bucketNum
		}
		buckets[bucketIndex] = append(buckets[bucketIndex], num)
	}
	for i := range buckets {
		QuickSort(buckets[i])
	}
	j := 0
	for _, bucket := range buckets {
		for _, num := range bucket {
			nums[j] = num
			j++
		}
	}
}

func RadixSort(nums []int) {
	var buckets [][]int
	max := math.MinInt64
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	stop := 1
	for ; max/stop != 0; stop *= 10 {
	}
	for base := 1; base != stop; base *= 10 {
		buckets = partitionBucket(nums, base)
		j := 0
		for _, bucket := range buckets {
			for _, num := range bucket {
				nums[j] = num
				j++
			}
		}
	}

}

func partitionBucket(nums []int, i int) [][]int {
	buckets := make([][]int, 10)
	for i := range buckets {
		buckets[i] = make([]int, 0)
	}
	base := i
	for _, num := range nums {
		bucketIndex := num % (base * 10) / base
		buckets[bucketIndex] = append(buckets[bucketIndex], num)
	}
	return buckets
}
