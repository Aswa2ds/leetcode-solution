package golearn

import (
	"bytes"
	"fmt"
	"runtime"
	"sort"
	"strconv"
	"sync"
	"testing"
	"time"
)

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func TestGoRoutineStoppedByVariable(t *testing.T) {
	running := true
	f1 := func() {
		for running {
			fmt.Printf("go routine: %d is running\n", GetGID())
		}
		fmt.Printf("go routine: %d terminated\n", GetGID())
	}
	go f1()
	go f1()
	time.Sleep(1 * time.Millisecond)
	running = false
}

func consumer(stop chan bool) {
	for {
		select {
		case <-stop:
			fmt.Printf("go routine: %d terminated\n", GetGID())
			return
		default:
			fmt.Printf("go routine: %d is running\n", GetGID())
		}
	}
}

func TestRoutineStoppedByChannel(t *testing.T) {
	stop := make(chan bool)
	waitGroup := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		go func() {
			waitGroup.Add(1)
			defer waitGroup.Done()
			consumer(stop)
		}()
	}
	waitForSignal()
	close(stop)
	fmt.Println("stop all go routine")
}

func waitForSignal() {
	time.Sleep(1 * time.Millisecond)
}

func TestSort(t *testing.T) {
	list := []int{2, 3, 1, 5, 4}
	sorted := make(chan bool)
	go func() { defer close(sorted); sort.Ints(list) }()
	<-sorted
	fmt.Println(list)
	time.Sleep(time.Millisecond)
	fmt.Println(list)
}

func TestBufferChannel(t *testing.T) {
	list := []int{2,1,3,5,4}
	sortIt := make(chan bool)
	go func() {
		sortIt<-true
		fmt.Println("hello world")
		sort.Ints(list)
	}()
	time.Sleep(time.Second)
	close(sortIt)
	time.Sleep(time.Millisecond)
	fmt.Println(list)
}

func TestForGoRoutine(t *testing.T) {
	var list = make([]int, 10)
	for i := range list {
		list[i] = i
	}
	for _, i := range list {
		i := i
		go func() {
			fmt.Print(i)
		}()
	}
	time.Sleep(time.Second)
}

type IntSlice []int

func (s *IntSlice) doSum(finished chan int) {
	sum := 0
	for _, num := range *s {
		sum += num
	}
	finished<-sum
}

func (s *IntSlice) splitToDoSum() int {
	sum := 0
	numCPU := runtime.NumCPU()
	finished := make(chan int, numCPU)
	start := time.Now()
	for i := 0; i < numCPU; i++ {
		go s.doSum(finished)
	}
	for i := 0; i < numCPU; i++ {
		sum += <-finished
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
	return sum
}

func TestParallelization(t *testing.T) {
	slice := make(IntSlice, 1000000000)
	for i := range slice {
		slice[i] = 1
	}
	// 1.086895649s
	fmt.Println(slice.splitToDoSum())
	//numCPU := runtime.NumCPU()
	//sum := 0
	//start := time.Now()
	//for i := 0; i < numCPU; i++ {
	//	for _, num := range slice {
	//		sum += num
	//	}
	//}
	//// 9.160649155s
	//end := time.Now()
	//fmt.Println(end.Sub(start))
	//fmt.Println(sum)

}

func TestMaxProcs(t *testing.T) {
	fmt.Println(runtime.GOMAXPROCS(0))
}