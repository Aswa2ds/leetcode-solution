package factory_method_pattern

type IFactory interface {
	create(owner string) Product
	createProduct(owner string) Product
	registerProduct(product Product)
}
