package main

import (
	"fmt"
	"strconv"
)

func main() {

	/*
		Coding Exercise #1
		Consider the following declarations:
		var i int = 3
		var f float64 = 3.2
		Write a Go program that converts i to float64 and f to int.
		Print out the type and the value of the variables created after conversion.
	*/

	var i int = 3
	var f float64 = 3.2

	f1 := float64(i)
	fmt.Printf("f1's type: %T, f1's value: %f\n", f1, f1)

	i1 := int(f)
	fmt.Printf("i1's type: %T, i1's value: %d\n", i1, i1)

	/*
		Coding Exercise #2
		Consider the following declarations:
		var i = 3
		var f = 3.2
		var s1, s2 = "3.14", "5"
		Write a Go program that converts:

		1. i to string (int to string)
		2. s2 to int (string to int)
		3. f to string (float64 to string)
		4. s1 to float64  (string to float64)
		5. Print the value and the type for each variable created after conversion.
	*/

	var i2 = 3
	var f2 = 3.2
	var s1, s2 = "3.14", "5"

	// 1. int to string
	s := strconv.Itoa(i2)
	fmt.Printf("s Type is %T, s value is %q\n", s, s)

	// 2. string to int
	is, err := strconv.Atoi(s2)
	if err == nil {
		fmt.Printf("i type is %T, i value is %v\n", is, is)
	} else {
		fmt.Println("Can not convert string to int.")
	}

	// 3. float64 to string
	ss1 := fmt.Sprintf("%f", f2)
	fmt.Printf("ss1's type: %T, ss1's value: %s\n", ss1, ss1)

	// 4. string to float64
	f1, err1 := strconv.ParseFloat(s1, 64)
	if err1 == nil {
		fmt.Printf("f1's type: %T, f1's value: %v\n", f1, f1)
	} else {
		fmt.Println("Cannot convert string to float64.")
	}

	/*
		Coding Exercise #3
		Create a Go program that computes how long does it take for the Sunlight to reach the Earth if we know that the distance from the Sun to Earth is 149.6 million km and the speed of light  is 299,792,458 m / s
		The formula used is: time = distance / speed
	*/

	const (
		distance = 149_600_000 * 1000 // distance from the Sun to Earth in m

		speed = 299_792_458 // speed of light in m / s
	)

	const time = distance / speed // time in seconds

	fmt.Printf("The Sunlight reaches Earth in %v seconds.\n", time)
}
