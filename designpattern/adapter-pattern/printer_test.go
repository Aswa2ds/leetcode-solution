package adapter_pattern

import "testing"

func TestOrderedPrinterAdapter_Print(t *testing.T) {
	printer := NewOrderedPrinter()
	printer.Print([]int{1,5,3,2,67,4,3,2})
}

func TestUnorderedPrinterAdapter_Print(t *testing.T) {
	printer := NewUnorderedPrinter()
	printer.Print([]int{1,5,3,2,67,4,3,2})
}
