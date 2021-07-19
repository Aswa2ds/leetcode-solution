package routinesolution

import (
	"fmt"
	"sync"
)

func printCat() {
	fmt.Println("cat")
}

func printDog() {
	fmt.Println("dog")
}

func printFish() {
	fmt.Println("fish")
}

// OrderedPrint channel method
func OrderedPrint() {
	var wg sync.WaitGroup
	var catCount uint64
	var dogCount uint64
	var fishCount uint64
	catCh := make(chan struct{}, 1)
	dogCh := make(chan struct{}, 1)
	fishCh := make(chan struct{}, 1)
	catCh <- struct{}{}

	wg.Add(3)
	go func() {
		for true {
			if catCount >= 100 {
				wg.Done()
				return
			}
			<- catCh
			fmt.Print(catCount, " ")
			printCat()
			catCount++
			dogCh <- struct{}{}
		}
	}()
	go func() {
		for true {
			if dogCount >= 100 {
				wg.Done()
				return
			}
			<- dogCh
			fmt.Print(dogCount, " ")
			printDog()
			dogCount++
			fishCh <- struct{}{}
		}
	}()
	go func() {
		for true {
			if fishCount >= 100 {
				wg.Done()
				return
			}
			<- fishCh
			fmt.Print(fishCount, " ")
			printFish()
			fishCount++
			catCh <- struct{}{}
		}
	}()

	wg.Wait()
}

//OrderedPrint mutex method
//func OrderedPrint() {
//	var catMutex, dogMutex, fishMutex sync.Mutex
//	dogMutex.Lock()
//	fishMutex.Lock()
//	go func() {
//		for i := 0; i < 100; i++ {
//			catMutex.Lock()
//			printCat()
//			dogMutex.Unlock()
//		}
//	}()
//	go func() {
//		for i := 0; i < 100; i++ {
//			dogMutex.Lock()
//			printDog()
//			fishMutex.Unlock()
//		}
//	}()
//	go func() {
//		for i := 0; i < 100; i++ {
//			fishMutex.Lock()
//			printFish()
//			catMutex.Unlock()
//		}
//	}()
//	time.Sleep(10*time.Second)
//}

