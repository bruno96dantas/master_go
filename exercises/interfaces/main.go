package main

import "fmt"

/*
Coding Exercise #1
1. Create a Go program where car type implements the vehicle interface.
2. Create a variable of type vehicle that holds a car struct value.
3. Call the methods (Licence and Name) on the interface value declared at step 2
*/
type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) License() string {
	return c.licenceNo
}

func (c car) Name() string {
	return c.brand
}

func main() {

	var v vehicle = car{brand: "Ford Mustang", licenceNo: "POW100ZZ"}

	fmt.Println(v.License())
	fmt.Println(v.Name())

	/*
		Coding Exercise #2
		1. Declare a variable called empty of type empty interface. Print out its type.
		2. Assign an int value to the variable called empty. Print out its type.
		3. Assign a float64 value to empty. Print out its type.
		4. Assign an int slice value to empty. Print out its type.
		5. Add a new int value to the slice (empty variable).
		6. Print out the slice (empty variable).
	*/

	var empty interface{}
	fmt.Printf("%T\n", empty)

	empty = 5
	fmt.Printf("%T\n", empty)

	empty = 10.10
	fmt.Printf("%T\n", empty)

	empty = []int{1, 2, 3}
	fmt.Printf("%T\n", empty)

	empty = append(empty.([]int), 10)
	fmt.Printf("%v\n", empty)
}
