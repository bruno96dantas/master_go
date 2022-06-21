package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Declaring two functions: f1 and f2
func f1(wg *sync.WaitGroup) { // wg is passed as a pointer
	fmt.Println("f1(goroutine) execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
		// sleep for a second to simulate an expensive task.
		time.Sleep(time.Second)

	}
	fmt.Println("f1 execution finished")

	//3.
	// Before exiting, call wg.Done() in each goroutine
	// to indicate to the WaitGroup that the goroutine has finished executing.
	wg.Done()
	//or:
	// (*wg).Done()
}

func f2() {
	fmt.Println("f2 execution started")
	time.Sleep(time.Second)

	for i := 5; i < 8; i++ {
		fmt.Println("f2(), i=", i)

	}
	fmt.Println("f2 execution finished")

}

func main() {
	fmt.Println("main execution started")

	var wg sync.WaitGroup

	wg.Add(1)

	go f1(&wg)

	fmt.Println("No. of Goroutines:", runtime.NumGoroutine()) // => 2

	f2()

	wg.Wait()

	fmt.Println("main execution stopped")
}
