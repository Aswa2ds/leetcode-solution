package adapter_pattern

type OrderedPrinter interface {
	Print([]int)
}

func NewOrderedPrinter() OrderedPrinter {
	return &OrderedPrinterAdapter{}
}
