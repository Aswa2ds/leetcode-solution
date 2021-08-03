package offer

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 || len(arr) == 0 {
		return make([]int, 0)
	}
	if k == len(arr) {
		return arr
	}
	return quickSearch(arr, 0, len(arr)-1, k)
}

func quickSearch(arr []int, low int, high int, k int) []int {
	j := partition(arr, low, high)
	switch true {
	case j > k:
		return quickSearch(arr, low, j-1, k)
	case j < k:
		return quickSearch(arr, j+1, high, k)
	default:
		return arr[:k]
	}
}

func partition(arr []int, low int, high int) int {
	pivot := arr[low]
	s1, s2 := low+1, high
	for true {
		for ; s1 < s2 && arr[s2] >= pivot; s2-- {}
		for ; s1 < s2 && arr[s1] <= pivot; s1++ {}
		if s1 != s2 {
			arr[s1], arr[s2] = arr[s2], arr[s1]
		} else {
			break
		}
	}
	if arr[s1] < arr[low] {
		arr[low], arr[s1] = arr[s1], arr[low]
		return s1
	} else {
		return low
	}
}

//heap method
//func getLeastNumbers(arr []int, k int) []int {
//	if k == 0 {
//		return make([]int, 0)
//	}
//	heap := make([]int, 0, k)
//	for _, num := range arr {
//		if len(heap) < cap(heap) {
//			heap = append(heap, num)
//			shiftUp(&heap, len(heap)-1)
//			continue
//		}
//		if heap[0] > num {
//			heap[0] = num
//			shiftDown(&heap, 0)
//		}
//	}
//	return heap
//}
//
//func shiftDown(heap *[]int, i int) {
//	left, right := 2*i+1, 2*i+2
//	max := i
//	if left >= len(*heap) {
//		return
//	}
//	if right >= len(*heap) {
//		if (*heap)[left] > (*heap)[max] {
//			max = left
//		}
//	} else {
//		if (*heap)[max] < (*heap)[left] {
//			max = left
//		}
//		if (*heap)[max] < (*heap)[right] {
//			max = right
//		}
//	}
//	if max != i {
//		(*heap)[i], (*heap)[max] = (*heap)[max], (*heap)[i]
//		shiftDown(heap, max)
//	}
//}
//
//func shiftUp(heap *[]int, i int) {
//	if i == 0 {
//		return
//	}
//	parent := (i-1)/2
//	if (*heap)[i] > (*heap)[parent] {
//		(*heap)[parent], (*heap)[i] = (*heap)[i], (*heap)[parent]
//		shiftUp(heap, parent)
//	}
//	return
//}
