package codetop

import (
	"fmt"
	"testing"
)

func TestOperation(t *testing.T) {
	lruCache := Constructor(2)
	lruCache.Put(2, 1)
	lruCache.Put(1, 1)
	lruCache.Put(2, 3)
	lruCache.Put(4, 1)
	fmt.Println(lruCache.Get(1))
	fmt.Println(lruCache.Get(2))
}
