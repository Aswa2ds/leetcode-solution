package iterator_pattern

import (
	"fmt"
	"testing"
)

func TestNewBookShelf(t *testing.T) {
	shelf := NewBookShelf([]Book{"bookA", "bookB", "bookC", "bookD"})
	for it := shelf.Iterator(); it.HasNext(); it.Next() {
		fmt.Println(it.GetItem())
	}
}
