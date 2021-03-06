package main

import (
	"fmt"
	"math"
	"sync"
)

/*
Coding Exercise #2
1. Create an anonymous function that calculates and prints out the square root of a float value it receives as argument.
2. Launch the function as a goroutine and synchronize it with main using WaitGroups
*/

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func(f float64, wg *sync.WaitGroup) {
		x := math.Sqrt(f)
		fmt.Printf("Square root of %.2f is %.4f\n", f, x)
		wg.Done()
	}(16.1, &wg)

	wg.Wait()
}
