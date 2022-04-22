package main

import "fmt"

type duration int

type mile float64
type kilometer float64

const m2km = 1.609

func main() {

	/*
		Coding Exercise #1
		1. Declare a new type called duration. Have the underlying type be an int.
		2. Declare a variable of the new type called hour using the var keyword
		3. In function main:
			a. print out the value of the variable hour
			b. print out the type of the variable hour
			c. assign 3600 to the variable hour using the  = operator
			d. print out the value of hour
	*/
	var hour duration
	fmt.Printf("hour's type: %T, hour's value: %v\n", hour, hour)

	hour = 3600
	fmt.Println(hour)

	/*
		Coding Exercise #2
		1. Declare two defined types called mile and kilometer. Have the underlying type be an float64.
		2. Declare a constant called m2km equals 1.609  ( 1mile=1.609km)
		3. In function main:
			a. declare a variable called mileBerlinToParis of type mile with value 655.3
			b. declare a variable called kmBerlinToParis of type kilometer
			c. calculate the distance between Berlin and Paris in km by multiplying mileBerlinToParis and m2km. Assign the result to kmBerlinToParis
			d. print out the distance in km between Berlin and Paris
	*/

	var mileBerlinToParis mile = 655.3 //distance in miles
	var kmBerlinToParis kilometer

	kmBerlinToParis = kilometer(mileBerlinToParis * m2km)
	fmt.Printf("Distance in Km from Berlin to Paris is %f\n", kmBerlinToParis)
}
