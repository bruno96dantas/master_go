package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		Coding Exercise #1
		Using a for loop and an if statement print out all the numbers between 1 and 50 that divisible by 7.
	*/

	for i := 1; i <= 50; i++ {
		if i%7 == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("")

	/*
		Coding Exercise #2
		Change the code from the previous exercise and use the continue statement to print out all the numbers divisible by 7 between 1 and 50.
	*/

	for i := 1; i <= 50; i++ {
		if i%7 != 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println("")

	/*
		Coding Exercise #3
		Change the code from the previous exercise and use the break statement to print out only the first 3 numbers divisible by 7 between 1 and 50.
	*/

	count := 0
	for i := 1; i <= 50; i++ {
		if i%7 != 0 {
			continue
		}
		fmt.Printf("%d ", i)
		count++

		if count == 3 {
			break
		}
	}
	fmt.Println("")

	/*
		Coding Exercise #4
		Using a for loop, an if statement and the logical and operator print out all the numbers between 1 and 500 that divisible both by 7 and 5.
	*/

	for i := 1; i <= 500; i++ {
		if i%5 == 0 && i%7 == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("")

	/*
		Coding Exercise #5
		Using a for loop print out all the years from your birthday to the current year.
		Use a variant of for loop where the post statement is moved inside the for block of code.
	*/

	birthYear := 1996
	currentYear := time.Now().Year()

	for i := birthYear; i <= currentYear; {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()
}
