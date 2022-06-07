package main

import "fmt"

func swap(a, b *float64) {
	*a, *b = *b, *a
}

func main() {

	/*
		Coding Exercise #1
		Consider the following variable declaration x := 10.10
		1. Print out the address of x
		2. Declare a pointer called ptr that stores the address of x.
		3. Print out the type and the value of ptr.
		4. Print the address of the pointer and the value of x though the pointer (use the dereferencing operator).
	*/

	x := 10.10

	// 1.
	fmt.Println("address:", &x)

	// 2.
	ptr := &x

	// 3.
	fmt.Printf("type: %T, value: %v\n", ptr, ptr)

	// 4.
	fmt.Println("address of ptr:", &ptr, "value of ptr:", *ptr)

	/*
		Coding Exercise #2
		Consider the following variable declaration:y, z := 5.5, 8.8
		Write a function that swaps the values of x and y. After calling the function x will be 8.8 and y will 5.5
	*/
	y, z := 5.5, 8.8
	swap(&y, &z)

	fmt.Printf("x is %v and y is %v\n", y, z)

}
