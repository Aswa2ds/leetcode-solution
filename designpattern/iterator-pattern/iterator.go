package iterator_pattern

type Iterator interface {
	GetItem() interface{}
	Next()
	HasNext() bool
}
