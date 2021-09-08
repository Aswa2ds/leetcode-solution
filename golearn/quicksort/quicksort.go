package quicksort

type SortInterface interface {
	Compare(int, int) bool
	Swap(int, int)
	Len() int
}

func Sort(data SortInterface) {
	QuickSort(data, 0, data.Len())
}

func QuickSort(data SortInterface, low, high int) {
	if low >= high {
		return
	}
	pivot := low
	a, b := low, high - 1
	for a < b {
		for ; a < b && data.Compare(pivot, b); b-- {}
		for ; a < b && !data.Compare(pivot, a); a++ {}
		data.Swap(a, b)
	}
	data.Swap(a, pivot)
	pivot = a
	QuickSort(data, low, pivot)
	QuickSort(data, pivot+1, high)
}