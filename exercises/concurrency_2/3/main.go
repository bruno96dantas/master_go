package main

import "fmt"

/*
Coding Exercise #3
Create a goroutine named power() that has one parameter of type int, calculates the square value of its parameter and then sends  the result into a channel.
In the main function launch 50 goroutines that calculate the square values of all numbers between 1 and 50 included.
Print out the square values.
*/
func power(x int, c chan int) {
	c <- x * x
}

func main() {
	ch := make(chan int)

	for i := 1; i <= 50; i++ {
		go power(i, ch)
	}

	for i := 1; i <= 50; i++ {
		fmt.Println(<-ch)
	}

}
