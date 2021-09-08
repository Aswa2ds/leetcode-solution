package iterator_pattern

type BookshelfIterator struct {
	bookShelf BookShelf
	index     int
}

func NewBookshelfIterator(bookShelf BookShelf) *BookshelfIterator {
	return &BookshelfIterator{bookShelf: bookShelf, index: 0}
}

func (it *BookshelfIterator) Next() {
	it.index++
}

func (it *BookshelfIterator) HasNext() bool {
	return it.index < it.bookShelf.Count()
}

func (it *BookshelfIterator) GetItem() interface{} {
	return it.bookShelf.GetBook(it.index)
}
