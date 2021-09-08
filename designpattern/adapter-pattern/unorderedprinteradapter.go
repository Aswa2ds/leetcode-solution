package adapter_pattern

type UnorderedPrinterAdapter struct {
	Printer
}

func (p *UnorderedPrinterAdapter) Print(list []int) {
	p.PrintList(list)
}