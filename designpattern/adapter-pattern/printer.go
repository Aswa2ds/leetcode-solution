package adapter_pattern

import "fmt"

type Printer struct {
}

func (p *Printer) PrintList(list []int) {
	for _, num := range list {
		fmt.Println(num)
	}
}