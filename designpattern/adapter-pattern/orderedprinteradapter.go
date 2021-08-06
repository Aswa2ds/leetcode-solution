package adapter_pattern

import "sort"

type BannerAdapter struct {
	Banner
}

func (b *BannerAdapter) PrintOrdered(list []int) {
	sort.Ints(list)
	b.PrintList(list)
}