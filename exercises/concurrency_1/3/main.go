package main

import (
	"fmt"
	"math"
	"sync"
)

/*
Coding Exercise #3
Change the code from Exercise #2 and launch 50 goroutines that calculate concurrently the square root of all the numbers between 100 and 149 (both included).
*/

func main() {
	var wg sync.WaitGroup

	wg.Add(50)
	for i := 100.; i < 150.; i++ {
		go func(f float64, wg *sync.WaitGroup) {
			x := math.Sqrt(f)
			fmt.Printf("Square root of %.2f is %.6f\n", f, x)
			wg.Done()
		}(i, &wg)
	}

	wg.Wait()
}
