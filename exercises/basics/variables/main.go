package main

import "fmt"

func main() {

	/*
		Coding Exercise #1
		1. Using the var keyword, declare 4 variables called a, b, c, d of type int, float64, bool and string
	*/
	var a int
	var b float64
	var c bool
	var d string

	/*
		2. Using short declaration syntax declare and assign these values to variables x, y and z:
		- 20
		- 15.5
		- "Gopher!"
	*/
	x := 20
	y := 15.5
	z := "Gopher!"

	fmt.Println(a, b, c, d)
	fmt.Println("x is", x, "y is", y, "z is", z)

	/*
		Coding Exercise #2
		1. Declare a, b, c, d using a single var keyword (multiple variable declaration) for better readability
	*/
	var (
		a2 int
		b2 float64
		c2 bool
		d2 string
	)

	/*
		2. Declare x, y, z on a single line -> multiple short declarations
	*/

	x2, y2, z2 := 20, 15.5, "Gopher!"

	/*
		3. Change the program to run without error using the blank identifier ( _ )
	*/

	_, _, _, _, _, _, _ = a2, b2, c2, d2, x2, y2, z2
}
