package offer

//func minArray(numbers []int) int {
//	min := numbers[0]
//	for _, number := range numbers {
//		if number < min {
//			return number
//		}
//	}
//	return min
//}

func minArray(numbers []int) int {
	low, high := 0, len(numbers)-1
	for low < high {
		pivot := (low + high) / 2
		if numbers[pivot] < numbers[high] {
			high = pivot
		} else if numbers[pivot] > numbers[high] {
			low = pivot + 1
		} else {
			high--
		}

	}
	return numbers[low]
}
