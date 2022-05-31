package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Coding Exercise #1
Create a function called cube() that takes a parameter of type float64 and returns the cube of that parameter (the parameter to the power of 3).
*/
func cube(n float64) float64 {

	return n * n * n
}

/*
Coding Exercise #2
Create a Go program with a function called f1() that takes a parameter of type uint and returns 2 values:
a) the factorial of n
b) the sum of all integer numbers greater than zero (>0) and less than or equal to n (<=n)
*/
func f1(n uint) (int, int) {

	fact := 1
	sum := 0

	for i := 1; i <= int(n); i++ {
		fact *= i
		sum += i
	}

	return fact, sum
}

/*
Coding Exercise #3
Write a function called myFunc() that takes exactly one argument which is an int number written between double quotes (this is in fact a string).
If the argument is integer 'n', the function should return the result of n + nn + nnn
*/
func myFunc(n string) int {

	i, err1 := strconv.Atoi(n)

	if err1 != nil {
		fmt.Printf("Cannot convert %q to int.", n)
		os.Exit(1)
	}
	ii, _ := strconv.Atoi(n + n)
	iii, _ := strconv.Atoi(n + n + n)

	return i + ii + iii
}

/*
Coding Exercise #4
Create a function with the identifier sum that takes in a variadic parameter of type int and returns the sum of all values of type int passed in.
*/
func sum(a ...int) int {
	s := 0
	for _, v := range a {
		s += v
	}

	return s
}

/*
Coding Exercise #5
Change the function from the previous exercise and use a `naked return`.
*/
func sum1(a ...int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}

/*
Coding Exercise #6
Create a function called searchItem() that takes 2 parameters: a) a string slice and b) a string.
The function should search for the string (the second parameter) in the slice (the first parameter) and returns true if it finds the string in the slice and false otherwise.
Do function does a case-sensitive search.
*/
func searchItem(mySlice []string, myStr string) bool {
	for _, v := range mySlice {
		if v == myStr {
			return true
		}

	}
	return false
}

/*
Coding Exercise #7
Change the function from the previous exercise to do a case-insensitive search.
*/
func searchItem1(mySlice []string, myStr string) bool {
	for _, v := range mySlice {
		if strings.EqualFold(v, myStr) {
			return true
		}

	}
	return false
}

func main() {

	// 1.
	v1 := cube(3.)
	fmt.Println("cube:", v1)

	// 2.
	v2, v3 := f1(5)
	fmt.Println("fact:", v2, "sum:", v3)

	// 3.
	fmt.Println(myFunc("5")) // => 615
	fmt.Println(myFunc("3")) // => 369

	// 4.
	s := sum(1, 2, 30)
	fmt.Println(s)

	// 5.
	s1 := sum1(1, 2, 30)
	fmt.Println(s1)

	// 6.
	animals := []string{"lion", "tiger", "bear"}

	result := searchItem(animals, "bear")
	fmt.Println(result) // => true

	result = searchItem(animals, "pig")
	fmt.Println(result) // => false

	// 7.
	result1 := searchItem1(animals, "BEAR")
	fmt.Println(result1) // => true

	result1 = searchItem1(animals, "lION")
	fmt.Println(result1) // => true

	result1 = searchItem1(animals, "pig")
	fmt.Println(result1) // => false

}
