package factory_method_pattern

import (
	"fmt"
)

type IDCardFactory struct {
	Factory
}

func (f *IDCardFactory) createProduct(owner string) Product {
	return &IDCard{owner: owner}
}

func (f *IDCardFactory) registerProduct(product Product) {
	fmt.Println(product.(*IDCard).getOwner() + " create IDCard")
}
