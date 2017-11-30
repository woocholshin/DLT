package myPackage

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	wg sync.WaitGroup

	mutex sync.Mutex
)

func Concurrent_test() string {
	/*
		runtime.GOMAXPROCS(runtime.NumCPU())

		runtime.Goshed() // give up control and let it go back to queue

	*/
	runtime.GOMAXPROCS(1)
	wg.Add(2)

	fmt.Println("#########################################")
	fmt.Println("running Go routine..")
	go printPrime("A")
	go printPrime("B")

	fmt.Println("waiting for Go routines to finish..")
	wg.Wait()

	return "Returned termination..."
}

func printPrime(prefix string) {
	defer wg.Done()

	mutex.Lock()
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("complete :", prefix)

	mutex.Unlock()
}
