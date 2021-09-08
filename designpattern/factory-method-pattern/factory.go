package factory_method_pattern

type Factory struct {
	Specific IFactory
}

func (f *Factory) create(owner string) Product {
	product := f.createProduct(owner)
	f.registerProduct(product)
	return product
}

func (f *Factory) createProduct(owner string) Product {
	if f.Specific == nil {
		return nil
	}
	return f.Specific.createProduct(owner)
}

func (f *Factory) registerProduct(product Product) {
	if f.Specific == nil {
		return
	}
	f.Specific.registerProduct(product)
}
