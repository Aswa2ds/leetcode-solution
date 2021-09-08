package iterator_pattern

type Book string

type BookShelf struct {
	books []Book
}

func NewBookShelf(books []Book) *BookShelf {
	return &BookShelf{books: books}
}

func (bs *BookShelf) Iterator() Iterator {
	return NewBookshelfIterator(*bs)
}

func (bs *BookShelf) Count() int {
	return len(bs.books)
}

func (bs *BookShelf) GetBook(index int) Book {
	if index >= len(bs.books) {
		return ""
	}
	return bs.books[index]
}

