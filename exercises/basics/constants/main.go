package main

import "fmt"

func main() {
	/*
		Coding Exercise #1
		Using the const keyword declare and initialize the following constants:
			1. daysWeek with value 7
			2. lightSpeed with value 299792458
			3. pi with value 3.14159
	*/

	const daysWeek int = 7
	const lightSpeed float64 = 299792458
	const pi float64 = 3.14159

	/*
		Coding Exercise #2
		Change the code from the previous exercise and declare all 3 constants as grouped constants.
		Make them untyped
	*/

	const (
		daysWeek1   = 7
		lightSpeed1 = 299792458
		pi1         = 3.14159
	)

	/*
		Coding Exercise #3
		Calculate how many seconds are in a year
		1. Declare secPerDay constant and initialize it to number of seconds in a day
		2. Declare daysYear constant and initialize it to 365
		3. Use PrintF to print out the total number of seconds in a year
	*/

	const (
		secPerDay = 60 * 60 * 24
		daysYear  = 365
	)

	fmt.Printf("There are %d seconds in a year. \n", secPerDay*daysYear)

	/*
		Coding Exercise #4
		Using Iota declare the following months of the year: Jun, Jul and Aug.
		Jun, Jul and Aug are constant and their value is 6, 7 and 8.
	*/

	const (
		Jun = iota + 6
		Jul
		Aug
	)

	fmt.Println(Jun, Jul, Aug)
}
