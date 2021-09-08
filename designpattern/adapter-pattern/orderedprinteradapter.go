package adapter_pattern

import "sort"

type OrderedPrinterAdapter struct {
	Printer
}

func (p *OrderedPrinterAdapter) Print(list []int) {
	sort.Ints(list)
	p.PrintList(list)
}