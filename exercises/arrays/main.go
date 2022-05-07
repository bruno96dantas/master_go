package main

import "fmt"

func main() {

	/*
		Coding Exercise #1
		1. Using the var keyword, declare an array called cities with 2 elements of type string. Don't initialize the array.
		Print out the cities array and notice the value of its elements

		2. Declare an array called grades of type [3]float64 and initialize only the first 2 elements using an array literal.
		Print out the cities array and notice the value of its elements

		3. Declare an array called taskDone using the ellipsis operator. The elements are of type bool.

		4. Iterate over the array called cities using the classical for loop syntax and len function and print out element by element.

		5. Iterate over grades using the range keyword and print out element by element.
	*/

	// 1.
	var cities = [2]string{}
	fmt.Printf("%#v\n", cities)

	// 2.
	var grades = [3]float64{2.2, 4.4}
	fmt.Printf("%#v\n", grades)

	// 3.
	taskDone := [...]bool{true, false, false, true}
	fmt.Printf("%#v\n", taskDone)

	// 4.
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	// 5.
	for index, value := range grades {
		fmt.Println(index, value)
	}

	/*
		Coding Exercise #2
		Consider the following array declaration: nums := [...]int{30, -1, -6, 90, -6}
		Write a Go program that counts the number of positive even numbers in the array.
	*/

	nums := [...]int{30, -1, -6, 90, -6}

	var count int

	for _, v := range nums {
		if v%2 == 0 && v > 0 {
			count++
		}
	}

	fmt.Println("No. of positive even numbers in nums: ", count)
}
