package iterator_pattern

type Aggregate interface {
	Iterator() Iterator
}
