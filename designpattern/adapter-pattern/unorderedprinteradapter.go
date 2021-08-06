package adapter_pattern

import "sort"

type OrderedPrinterAdapter struct {
	Printer
}

func (p *OrderedPrinterAdapter) PrintOrdered(list []int) {
	sort.Ints(list)
	p.PrintList(list)
}