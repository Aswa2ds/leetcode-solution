package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func Test_GoRoutine(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Println("hello, world")
		wg.Done()
	}()
	wg.Wait()
}
