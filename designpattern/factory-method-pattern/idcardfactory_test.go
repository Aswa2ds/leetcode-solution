package factory_method_pattern

import (
	"testing"
)

func TestIDCardFactory(t *testing.T) {
	f := &Factory{Specific: &IDCardFactory{}}
	product := f.create("Mike")
	product.use()
}
