package basicalgorithm

import (
	"math/rand"
	"testing"
	"time"
)

type SortedError struct {

}

func (e *SortedError) Error() string {
	return "sorted error"
}

func isSorted(list []int) bool {
	if len(list) < 1 {
		return true
	}
	for i := 0; i < len(list)-1; i++ {
		if list[i] > list[i+1]  {
			return false
		}
	}
	return true
}

func generateRandomList(n int) []int {
	list := make([]int, n)
	rand.Seed(time.Now().UnixNano())
	for i := range list {
		list[i] = rand.Intn(1000000)
	}
	return list
}

func TestBubbleSort(t *testing.T) {
	n := 100000
	list := generateRandomList(n)
	// list := []int{4,7,1,0,8,3,2,9,5,6}
	// fmt.Println(list)
	BubbleSort(list)
	// fmt.Println(list)
	if !isSorted(list) {
		t.Error(new(SortedError))
	}
}

func TestBubbleSortWithFlag(t *testing.T) {
	n := 100000
	list := generateRandomList(n)
	// list := []int{4,7,1,0,8,3,2,9,5,6}
	// fmt.Println(list)
	BubbleSortWithFlag(list)
	// fmt.Println(list)
	if !isSorted(list) {
		t.Error(new(SortedError))
	}
}

func TestBubbleSortWithRecord(t *testing.T) {
	n := 100000
	list := generateRandomList(n)
	// list := []int{4,7,1,0,8,3,2,9,5,6}
	// fmt.Println(list)
	BubbleSortWithRecord(list)
	// fmt.Println(list)
	if !isSorted(list) {
		t.Error(new(SortedError))
	}
}

func TestCockTailSort(t *testing.T) {
	n := 100000
	list := generateRandomList(n)
	// list := []int{4,7,1,0,8,3,2,9,5,6}
	// fmt.Println(list)
	CockTailSort(list)
	// fmt.Println(list)
	if !isSorted(list) {
		t.Error(new(SortedError))
	}
}

func TestQuickSort(t *testing.T) {
	n := 100000
	list := generateRandomList(n)
	// list := []int{4,7,1,0,8,3,2,9,5,6}
	// fmt.Println(list)
	QuickSort(list)
	// fmt.Println(list)
	if !isSorted(list) {
		t.Error(new(SortedError))
	}
}

func TestQuickSortNoRecursion(t *testing.T) {
	n := 100000
	list := generateRandomList(n)
	// list := []int{4,7,1,0,8,3,2,9,5,6}
	// fmt.Println(list)
	QuickSortNoRecursion(list)
	// fmt.Println(list)
	if !isSorted(list) {
		t.Error(new(SortedError))
	}
}

func TestHeapSort(t *testing.T) {
	n := 100000
	list := generateRandomList(n)
	// list := []int{4,7,1,0,8,3,2,9,5,6}
	// fmt.Println(list)
	HeapSort(list)
	// fmt.Println(list)
	if !isSorted(list) {
		t.Error(new(SortedError))
	}
}

func TestSelectSort(t *testing.T) {
	n := 100000
	list := generateRandomList(n)
	// list := []int{4,7,1,0,8,3,2,9,5,6}
	// fmt.Println(list)
	SelectSort(list)
	// fmt.Println(list)
	if !isSorted(list) {
		t.Error(new(SortedError))
	}
}

func TestInsertSort(t *testing.T) {
	n := 100000
	list := generateRandomList(n)
	// list := []int{4,7,1,0,8,3,2,9,5,6}
	// fmt.Println(list)
	InsertSort(list)
	// fmt.Println(list)
	if !isSorted(list) {
		t.Error(new(SortedError))
	}
}

func TestInsertSortWithQuickSearch(t *testing.T) {
	n := 100000
	list := generateRandomList(n)
	// list := []int{4,7,1,0,8,3,2,9,5,6}
	// fmt.Println(list)
	InsertSortWithQuickSearch(list)
	// fmt.Println(list)
	if !isSorted(list) {
		t.Error(new(SortedError))
	}
}

func TestShellSort(t *testing.T) {
	n := 100000
	list := generateRandomList(n)
	// list := []int{4,7,1,0,8,3,2,9,5,6}
	// fmt.Println(list)
	ShellSort(list)
	// fmt.Println(list)
	if !isSorted(list) {
		t.Error(new(SortedError))
	}
}

func TestMergeSort(t *testing.T) {
	n := 100000
	list := generateRandomList(n)
	// list := []int{4,7,1,0,8,3,2,9,5,6}
	// fmt.Println(list)
	list = MergeSort(list)
	// fmt.Println(list)
	if !isSorted(list) {
		t.Error(new(SortedError))
	}
}

func TestCountSort(t *testing.T) {
	n := 100000
	list := generateRandomList(n)
	// list := []int{4,7,1,0,8,3,2,9,5,6}
	// fmt.Println(list)
	CountSort(list)
	// fmt.Println(list)
	if !isSorted(list) {
		t.Error(new(SortedError))
	}
}

func TestBucketSort(t *testing.T) {
	n := 100000
	list := generateRandomList(n)
	// list := []int{4,7,1,0,8,3,2,9,5,6}
	// fmt.Println(list)
	BucketSort(list)
	// fmt.Println(list)
	if !isSorted(list) {
		t.Error(new(SortedError))
	}
}

func TestRadixSort(t *testing.T) {
	n := 100000
	list := generateRandomList(n)
	// list := []int{4,7,1,0,8,3,2,9,5,6}
	// fmt.Println(list)
	RadixSort(list)
	// fmt.Println(list)
	if !isSorted(list) {
		t.Error(new(SortedError))
	}
}