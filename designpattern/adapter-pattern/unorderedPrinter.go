package adapter_pattern

type UnorderedPrinter interface {
	Print([]int)
}

func NewUnorderedPrinter() UnorderedPrinter {
	return &UnorderedPrinterAdapter{}
}
