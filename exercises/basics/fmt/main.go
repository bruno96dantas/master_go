package main

import "fmt"

func main() {

	/*
		Coding Exercise #1
		1. Print each variable using a specific verb for its type
		2. Print the string value enclosed by double quotes ("Gophers")
		3. Print each variable using the same verb
		4. Print the type of y and score
	*/
	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}

	// 1.
	fmt.Printf("x is %d, y is %f, z is %s\n", x, y, z)
	fmt.Printf("score is %#v\n", score)

	// 2.
	fmt.Printf("z is %q\n", z)

	// 3.
	fmt.Printf("x is %v, y is %v, z is %v, score is %v\n", x, y, z, score)

	// 4.
	fmt.Printf("y type: %T, score type: %T\n", y, score)

	/*
		Coding Exercise #2
		Consider the following constant declaration: const x float64 = 1.422349587101
		Write a Go program that prints the value of x with 4 decimal points.
	*/

	const x1 float64 = 1.422349587101

	fmt.Printf("x1 is %.4f\n", x1)
}
