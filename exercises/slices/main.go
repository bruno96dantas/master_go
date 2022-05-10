package main

import (
	"fmt"
)

func main() {

	/*
		Coding Exercise #1
		Using a composite literal declare and initialize a slice of type string with 3 elements.
		Iterate over the slice and print each element in the slice and its index.
	*/

	str := []string{"a", "b", "c"}
	for i, v := range str {
		fmt.Printf("Index: %d, Value: %q\n", i, v)
	}

	/*
		Coding Exercise #2
		1. Declare a slice called nums with three float64 numbers.
		2. Append the value 10.1 to the slice
		3. In one statement append to the slice the values: 4.1,  5.5 and 6.6
		4. Print out the slice
		5. Declare a slice called n with two float64 values
		6. Append n to nums
		7. Print out the nums slice
	*/

	nums := []float64{1.1, 2.2, 3.3}
	nums = append(nums, 10.1)
	nums = append(nums, 4.1, 5.5, 6.6)

	fmt.Println(nums)

	n := []float64{12.3, 13.4}
	n = append(n, nums...)

	fmt.Println(n)

	/*
		Coding Exercise #3
		Consider the following slice declaration: nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
		Using a slice expression and a for loop iterate over the slice ignoring the first and the last two elements.
		Print those elements and their sum.
	*/

	nums1 := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	sum := 0
	for _, v := range nums1[2 : len(nums1)-2] {
		fmt.Println(v)
		sum += v
	}
	fmt.Println("Sum:", sum)

	/*
		Coding Exercise #4
		Consider the following slice declaration: friends := []string{"Marry", "John", "Paul", "Diana"}
		Using copy() function create a copy of the slice. Prove that the slices are not connected by modifying one slice and notice that the other slice is not modified.
	*/

	friends := []string{"Marry", "John", "Paul", "Diana"}
	yourFriends := make([]string, len(friends))
	copy(yourFriends, friends)

	yourFriends[0] = "Dan"

	fmt.Println(friends, yourFriends)

	/*
		Coding Exercise #5
		Consider the following slice declaration: friends := []string{"Marry", "John", "Paul", "Diana"}
		Using append() function create a copy of the slice. Prove that the slices are not connected by modifying one slice and notice that the other slice is not modified.
	*/

	var friends2 []string
	friends2 = append(friends2, friends...)

	friends2[0] = "Dan"

	fmt.Println(friends, friends2)

	/*
		Coding Exercise #6
		Consider the following slice declaration:
		 years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
		Using a slice expression and append() function create a new slice called newYears that contains the first 3 and the last 3 elements of the slice.  newYears should be []int{2000, 2001, 2002, 2008, 2009, 2010}
	*/

	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	var newYears []int

	newYears = append(years[:3], years[len(years)-3:]...)

	fmt.Printf("%#v\n", newYears)
}
