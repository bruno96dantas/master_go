package main

import "fmt"

/*
Coding Exercise #1
1. Create a new type called money. Its underlying type is float64.
2. Create a method called print that prints out the money value with exact 2 decimal points.
*/

type money float64

func (m money) print() {
	fmt.Printf("%.2f\n", m)
}

/*
Coding Exercise #2
Consider the money type declared at exercise #1. Create a new method for the money type called printStr that returns the money value as a string with 2 decimal points.
*/
func (m money) printStr() string {
	return fmt.Sprintf("%.2f\n", m)
}

/*
Coding Exercise #3
1. Create a new struct type called book with 2 fields: price and title of type float64 and string.
2. Create a method for the newly defined type called vat that returns the vat value if vat is 9%.
3. Create a method called discount that takes a book value as receiver and decreases the price of the book by 10%.
*/
type book struct {
	title string
	price float64
}

func (b book) vat() float64 {
	return b.price * 0.09
}

func (b *book) discount() {
	(*b).price = b.price * 0.9
}

func main() {
	// 1.
	var m money = 45.6 * 879
	m.print()

	// 2.
	s := m.printStr()
	fmt.Println(s)

	// 3.
	bestBook := book{title: "The trial", price: 9.9}

	vat := bestBook.vat()
	fmt.Printf("Vat:%v\n", vat)

	bestBook.discount()
	fmt.Printf("%#v\n", bestBook)
}
