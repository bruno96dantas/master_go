package main

import "fmt"

// 1.
type person struct {
	name   string
	age    int
	colors []string
}

func main() {

	/*
		Coding Exercise #1
		1. Create your own struct type called person that stores the following data: name, age and a list with favorite colors.
		2. Declare and initialize two values of type person, one called me and another called you.
		3. Print out the struct values.
	*/

	// 2.
	me := person{name: "Marius", age: 30, colors: []string{"red", "yellow"}}
	you := person{name: "Maria", age: 22, colors: []string{"pink", "blue"}}

	// 3.
	fmt.Printf("%v\n", me)
	fmt.Printf("%+v\n", you)

	/*
		Coding Exercise #2
		Consider the code from the previous exercise and:
		1. Change the name or the struct value called me to "Andrei"
		2. Take in a new variable the favorites colors of struct value called you. Print out the type and the value of the new variable.
		3. Add a new favorite color to the second struct value called you.
		4. Print out the struct values.
	*/

	// 1.
	me.name = "Andrei"

	// 2.
	var colors []string = you.colors
	fmt.Printf("Type: %T, Value: %v\n", colors, colors)

	// 3.
	colors = append(colors, "green")
	you.colors = colors

	// 4.
	fmt.Printf("%v\n", me)
	fmt.Printf("%+v\n", you)

	/*
		Coding Exercise #3
		Consider the code from Exercise #1.
		Iterate and print out the favorite colors of the struct value called me.
	*/

	for index, color := range me.colors {
		fmt.Printf("%d -> %q\n", index, color)
	}

	/*
		Coding Exercise #4
		Change the code from Exercise #1 and:
		1. Create a new struct type called grades with  2 fields: grade and course
		2. Add another field of type grades to person struct type (embedded struct).
		3. Change the initialization of the struct values called me and you to contain information about the grades.
		4. Change the grade and the course of one struct value to "Golang" and 98.
		5. Print out the struct values.
	*/

	// 1.
	type grades struct {
		grade  int
		course string
	}

	// 2
	type person1 struct {
		name   string
		age    int
		colors []string
		gr     grades
	}

	// 3.
	me1 := person1{
		name:   "Marius",
		age:    30,
		colors: []string{"red", "yellow"},
		gr:     grades{grade: 85, course: "Python"},
	}
	you1 := person1{
		name:   "Maria",
		age:    22,
		colors: []string{"pink", "blue"},
		gr:     grades{grade: 100, course: "History"},
	}

	// 4.
	me1.gr.grade = 98
	me1.gr.course = "Golang"

	// 5.
	fmt.Printf("%v\n", me1)
	fmt.Printf("%+v\n", you1)

}
